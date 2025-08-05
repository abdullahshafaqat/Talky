// websocket/hub.go
package websocket

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte
}

type Hub struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
	mu         sync.RWMutex
}

var (
	hub  *Hub
	once sync.Once
)

// InitHub initializes the WebSocket hub (call this once at startup)
func InitHub() *Hub {
	once.Do(func() {
		hub = &Hub{
			Clients:    make(map[string]*Client),
			Register:   make(chan *Client),
			Unregister: make(chan *Client),
			Broadcast:  make(chan []byte),
		}
		go hub.run()
		log.Println("✅ [Hub] WebSocket hub initialized and running")
	})
	return hub
}

// GetHub returns the singleton hub instance
func GetHub() *Hub {
	if hub == nil {
		log.Println("⚠️ [Hub] Hub not initialized, initializing now...")
		return InitHub()
	}
	return hub
}

func (h *Hub) run() {
	log.Println("🚀 [Hub] Starting hub event loop...")
	
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client.ID] = client
			h.mu.Unlock()
			log.Printf("✅ [Hub] Client connected: %s (Total: %d)", client.ID, len(h.Clients))

		case client := <-h.Unregister:
			h.mu.Lock()
			if _, ok := h.Clients[client.ID]; ok {
				delete(h.Clients, client.ID)
				close(client.Send)
				log.Printf("❌ [Hub] Client disconnected: %s (Total: %d)", client.ID, len(h.Clients))
			}
			h.mu.Unlock()

		case message := <-h.Broadcast:
			h.mu.RLock()
			clientCount := len(h.Clients)
			log.Printf("📡 [Hub] Broadcasting message to %d clients: %s", clientCount, string(message))
			
			for clientID, client := range h.Clients {
				select {
				case client.Send <- message:
					log.Printf("📤 [Hub] Message sent to client: %s", clientID)
				default:
					// Client's send channel is full, remove client
					log.Printf("⚠️ [Hub] Client %s send channel full, removing client", clientID)
					close(client.Send)
					delete(h.Clients, clientID)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// GetConnectedClients returns list of connected client IDs
func (h *Hub) GetConnectedClients() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	
	clients := make([]string, 0, len(h.Clients))
	for clientID := range h.Clients {
		clients = append(clients, clientID)
	}
	return clients
}

// SendToUser sends a message to a specific user
func (h *Hub) SendToUser(userID string, message []byte) bool {
	h.mu.RLock()
	client, exists := h.Clients[userID]
	h.mu.RUnlock()
	
	if !exists {
		log.Printf("⚠️ [Hub] User %s not connected", userID)
		return false
	}
	
	select {
	case client.Send <- message:
		log.Printf("📤 [Hub] Direct message sent to user: %s", userID)
		return true
	default:
		log.Printf("⚠️ [Hub] Failed to send message to user %s (channel full)", userID)
		return false
	}
}