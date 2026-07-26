package p2p

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRoom_otherPeer(t *testing.T) {
	type args struct {
		c *Client
	}
	tests := []struct {
		name string
		r    *Room
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.otherPeer(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Room.otherPeer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_addPeer(t *testing.T) {
	type args struct {
		c *Client
	}
	tests := []struct {
		name string
		r    *Room
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.addPeer(tt.args.c); got != tt.want {
				t.Errorf("Room.addPeer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_removePeer(t *testing.T) {
	type args struct {
		c *Client
	}
	tests := []struct {
		name string
		r    *Room
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.removePeer(tt.args.c)
		})
	}
}

func TestRoom_isEmpty(t *testing.T) {
	tests := []struct {
		name string
		r    *Room
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.isEmpty(); got != tt.want {
				t.Errorf("Room.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHub(t *testing.T) {
	tests := []struct {
		name string
		want *Hub
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHub(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRoomID(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRoomID()
			if (err != nil) != tt.wantErr {
				t.Errorf("generateRoomID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateRoomID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHub_HandleCreateRoom(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *Hub
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.HandleCreateRoom(tt.args.w, tt.args.r)
		})
	}
}

func TestHub_HandleWebSocket(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *Hub
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.HandleWebSocket(tt.args.w, tt.args.r)
		})
	}
}

func TestClient_readPump(t *testing.T) {
	type args struct {
		h    *Hub
		room *Room
	}
	tests := []struct {
		name string
		c    *Client
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.readPump(tt.args.h, tt.args.room)
		})
	}
}

func TestClient_writePump(t *testing.T) {
	tests := []struct {
		name string
		c    *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.writePump()
		})
	}
}

func Test_mustMarshal(t *testing.T) {
	type args struct {
		e Envelope
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mustMarshal(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mustMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHub_CleanupLoop(t *testing.T) {
	tests := []struct {
		name string
		h    *Hub
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.CleanupLoop()
		})
	}
}
