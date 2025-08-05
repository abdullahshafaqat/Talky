package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	messageservice "github.com/abdullahshafaqat/Chatify/api/message_service"
	"github.com/abdullahshafaqat/Chatify/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (adjust in production)
	},
}

type MessageNotification struct {
	Type    string         `json:"type"`
	Message *models.Message `json:"message"`
}

var messageService messageservice.Service

func init() {
	// Initialize message service
	messageService = messageservice.NewMessageService()
}

func ServeWebSocket(c *gin.Context) {
	// Get user_id from middleware context with better error handling
	userID := c.GetString("user_id")
	if userID == "" {
		log.Println("❌ [WebSocket] user_id not found in context")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	log.Printf("🔌 [WebSocket] New connection attempt for user: %s", userID)

	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("❌ [WebSocket] WebSocket upgrade failed for user %s: %v", userID, err)
		return
	}

	log.Printf("✅ [WebSocket] Connection established for user: %s", userID)

	// Create client
	client := &Client{
		ID:   userID,
		Conn: conn,
		Send: make(chan []byte, 256), // Buffered channel to prevent blocking
	}

	// Get hub and register client
	hub := GetHub()
	hub.Register <- client

	// Start goroutines for reading and writing messages
	go readMessages(client)
	go writeMessages(client)
}

func readMessages(client *Client) {
	defer func() {
		log.Printf("🔌 [WebSocket] Closing read connection for user: %s", client.ID)
		GetHub().Unregister <- client
		client.Conn.Close()
	}()

	// Set read deadline and pong handler for connection health
	client.Conn.SetReadLimit(512)
	client.Conn.SetPongHandler(func(string) error {
		log.Printf("🏓 [WebSocket] Pong received from user: %s", client.ID)
		return nil
	})

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("❌ [WebSocket] Unexpected close error from %s: %v", client.ID, err)
			} else {
				log.Printf("📴 [WebSocket] Read error from %s: %v", client.ID, err)
			}
			break
		}

		log.Printf("📨 [WebSocket] Received from %s: %s", client.ID, string(msg))

		// Parse incoming message
		var notification MessageNotification
		if err := json.Unmarshal(msg, &notification); err != nil { // Fixed to use &notification
			log.Printf("❌ [WebSocket] Failed to parse notification for user %s: %v", client.ID, err)
			GetHub().Broadcast <- msg
			continue
		}

		hub := GetHub()
		switch notification.Type {
		case "message":
			if notification.Message != nil {
				notification.Message.SenderID = client.ID
				notification.Message.Timestamp = time.Now()
				data, err := json.Marshal(notification)
				if err != nil {
					log.Printf("❌ [WebSocket] Failed to marshal message for user %s: %v", client.ID, err)
					GetHub().Broadcast <- msg
					continue
				}
				// Broadcast and mark delivered for direct recipient
				if hub.SendToUser(notification.Message.ReceiverID, data) {
					if err := messageService.MarkMessageDelivered(notification.Message.ID, notification.Message.ReceiverID); err != nil {
						log.Printf("❌ [WebSocket] Failed to mark message %s as delivered: %v", notification.Message.ID, err)
					} else {
						notification.Message.Delivered = true
						// Notify sender of delivery
						deliveredNotification := &MessageNotification{
							Type:    "delivered",
							Message: notification.Message,
						}
						deliveredData, err := json.Marshal(deliveredNotification)
						if err != nil {
							log.Printf("❌ [WebSocket] Failed to marshal delivered notification: %v", err)
						} else {
							hub.SendToUser(client.ID, deliveredData)
						}
					}
				}
				// Broadcast to all except sender
				GetHub().Broadcast <- data
			}
		case "seen":
			if notification.Message != nil && notification.Message.ID != "" {
				if err := messageService.MarkMessageSeen(notification.Message.ID, client.ID); err != nil {
					log.Printf("❌ [WebSocket] Failed to mark message %s as seen for user %s: %v", notification.Message.ID, client.ID, err)
				} else {
					// Notify sender of seen status
					seenNotification := &MessageNotification{
						Type:    "seen",
						Message: notification.Message,
					}
					seenData, err := json.Marshal(seenNotification)
					if err != nil {
						log.Printf("❌ [WebSocket] Failed to marshal seen notification: %v", err)
					} else {
						hub.SendToUser(notification.Message.SenderID, seenData)
					}
				}
			}
		default:
			// Broadcast unprocessed messages
			GetHub().Broadcast <- msg
		}
	}
}

func writeMessages(client *Client) {
	defer func() {
		log.Printf("🔌 [WebSocket] Closing write connection for user: %s", client.ID)
		client.Conn.Close()
	}()

	// Use for range to iterate over the channel
	for msg := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("❌ [WebSocket] Write error to %s: %v", client.ID, err)
			return
		}

		log.Printf("📤 [WebSocket] Sent to %s: %s", client.ID, string(msg))
	}
}