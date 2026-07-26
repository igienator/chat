package p2p

import (
	"log"
	"net/http"
	"time"
)

func NewServer() *http.Server {
	hub := NewHub()
	go hub.CleanupLoop()

	mux := http.NewServeMux()
	mux.HandleFunc("/room", hub.HandleCreateRoom)
	mux.HandleFunc("/ws", hub.HandleWebSocket)

	return &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}

func StartP2P() error {
	srv := NewServer()
	log.Println("room-chat signaling/relay server listening on :8080")
	return srv.ListenAndServe()
}
