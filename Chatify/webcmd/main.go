package main

import (
	"log"
	"os"
	"github.com/abdullahshafaqat/Chatify/config"
	"github.com/abdullahshafaqat/Chatify/middleware"
	"github.com/abdullahshafaqat/Chatify/websocket"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("🚀 Starting Chatify WebSocket Server...")

	// Load configuration
	cfg := config.LoadConfig()
	log.Printf("✅ Configuration loaded. JWT Secret length: %d bytes", len(cfg.JWTSecret))

	// Initialize JWT secret
	middleware.InitAuthMiddleware()

	// Initialize WebSocket hub
	websocket.InitHub()

	// Create Gin router
	router := gin.Default()

	// Add CORS middleware for WebSocket connections
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Register WebSocket routes
	websocket.RegisterWebSocketRoutes(router)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8003"
	}

	log.Printf("🚀 Starting WebSocket server on :%s", port)
	log.Println("📡 WebSocket endpoint: ws://localhost:" + port + "/ws/connect?Authorization=your-jwt-token")
	log.Println("🔍 Health check: http://localhost:" + port + "/ws/health")
	log.Printf("🔑 MongoDB URI: %s", cfg.MongoURI)
	log.Printf("🔑 Firebase Key Path: %s", cfg.FirebaseKeyPath)
	
	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}