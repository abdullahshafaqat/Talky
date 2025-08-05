// websocket/routes.go
package websocket

import (
	"log"
	"net/http"

	"github.com/abdullahshafaqat/Chatify/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterWebSocketRoutes(router *gin.Engine) {
	log.Println("🚀 [Routes] Registering WebSocket routes...")

	// Create WebSocket group with auth middleware
	wsGroup := router.Group("/ws")
	wsGroup.Use(middleware.WebSocketAuthMiddleware())

	// WebSocket connection endpoint
	wsGroup.GET("/connect", ServeWebSocket)

	// Optional: Add a health check endpoint (without auth)
	router.GET("/ws/health", func(c *gin.Context) {
		hub := GetHub()
		clients := hub.GetConnectedClients()
		
		c.JSON(http.StatusOK, gin.H{
			"status":           "ok",
			"connected_clients": len(clients),
			"client_ids":       clients,
		})
	})

	log.Println("✅ [Routes] WebSocket routes registered:")
	log.Println("   - GET /ws/connect (WebSocket endpoint with auth)")
	log.Println("   - GET /ws/health (Health check without auth)")
}