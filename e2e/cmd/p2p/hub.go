package p2p

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	roomIDLength     = 6
	roomIDAlphabet   = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789" // no ambiguous chars (no 0/O, 1/I)
	unclaimedRoomTTL = 10 * time.Minute
	cleanupInterval  = time.Minute
)

// Envelope is the only structure the server understands on the wire.
// Payload is opaque to the server: it's the client's RSA public key,
// an AES-GCM encrypted chat message, or a control signal. The server
// never decrypts, inspects, or logs the payload contents -- only the
// "type" field, for routing and connection-health logging.
type Envelope struct {
	Type    string `json:"type"`    // "pubkey" | "message" | "leave"
	Payload string `json:"payload"` // meaningful only to the two clients
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
	room *Room
}

type Room struct {
	id      string
	peers   [2]*Client
	created time.Time
	mu      sync.Mutex
}

func (r *Room) otherPeer(c *Client) *Client {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, p := range r.peers {
		if p != nil && p != c {
			return p
		}
	}
	return nil
}

func (r *Room) addPeer(c *Client) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, p := range r.peers {
		if p == nil {
			r.peers[i] = c
			c.room = r
			return true
		}
	}
	return false // room already has two peers
}

func (r *Room) removePeer(c *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, p := range r.peers {
		if p == c {
			r.peers[i] = nil
		}
	}
}

func (r *Room) isEmpty() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.peers[0] == nil && r.peers[1] == nil
}

type Hub struct {
	mu       sync.Mutex
	rooms    map[string]*Room
	upgrader websocket.Upgrader
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[string]*Room),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true }, // native app, no browser origin to check
		},
	}
}

func generateRoomID() (string, error) {
	b := make([]byte, roomIDLength)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(roomIDAlphabet))))
		if err != nil {
			return "", err
		}
		b[i] = roomIDAlphabet[n.Int64()]
	}
	return string(b), nil
}

// HandleCreateRoom issues a fresh room ID. The room exists in memory only
// and expires automatically if nobody joins within unclaimedRoomTTL.
func (h *Hub) HandleCreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	var id string
	for {
		candidate, err := generateRoomID()
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		if _, exists := h.rooms[candidate]; !exists {
			id = candidate
			break
		}
	}

	h.rooms[id] = &Room{id: id, created: time.Now()}
	log.Printf("room created: %s", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"room_id": id})
}

// HandleWebSocket upgrades the connection and attaches the client to an
// existing room. Once two peers are present, every frame one sends is
// relayed verbatim to the other -- the server does not parse payloads.
func (h *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room")
	if roomID == "" {
		http.Error(w, "missing room id", http.StatusBadRequest)
		return
	}

	h.mu.Lock()
	room, ok := h.rooms[roomID]
	h.mu.Unlock()
	if !ok {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade failed: %v", err)
		return
	}

	client := &Client{conn: conn, send: make(chan []byte, 16)}
	if !room.addPeer(client) {
		conn.WriteJSON(Envelope{Type: "error", Payload: "room full"})
		conn.Close()
		return
	}

	log.Printf("peer joined room %s", roomID)

	if peer := room.otherPeer(client); peer != nil {
		peer.send <- mustMarshal(Envelope{Type: "peer-joined"})
		client.send <- mustMarshal(Envelope{Type: "peer-joined"})
	}

	go client.writePump()
	client.readPump(h, room)
}

func (c *Client) readPump(h *Hub, room *Room) {
	defer func() {
		room.removePeer(c)
		if peer := room.otherPeer(c); peer != nil {
			peer.send <- mustMarshal(Envelope{Type: "peer-left"})
		}
		close(c.send)
		c.conn.Close()

		if room.isEmpty() {
			h.mu.Lock()
			delete(h.rooms, room.id)
			h.mu.Unlock()
			log.Printf("room closed: %s", room.id)
		}
	}()

	c.conn.SetReadLimit(64 * 1024) // frames are ciphertext/keys, not file transfers

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		var env Envelope
		if err := json.Unmarshal(data, &env); err != nil {
			continue // drop malformed frames without killing the connection
		}

		// Log only the frame type for operational visibility -- never the payload.
		log.Printf("relaying %q frame in room %s", env.Type, room.id)

		if peer := room.otherPeer(c); peer != nil {
			peer.send <- data
		}
	}
}

func (c *Client) writePump() {
	for msg := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			return
		}
	}
}

func mustMarshal(e Envelope) []byte {
	b, _ := json.Marshal(e)
	return b
}

// CleanupLoop periodically removes rooms that were created but never
// claimed by a second peer within unclaimedRoomTTL.
func (h *Hub) CleanupLoop() {
	ticker := time.NewTicker(cleanupInterval)
	defer ticker.Stop()
	for range ticker.C {
		h.mu.Lock()
		for id, room := range h.rooms {
			if room.isEmpty() && time.Since(room.created) > unclaimedRoomTTL {
				delete(h.rooms, id)
				log.Printf("expired unclaimed room: %s", id)
			}
		}
		h.mu.Unlock()
	}
}
