package main

import (
	"log"
	"github.com/abdullahshafaqat/Chatify/config"
	"github.com/abdullahshafaqat/Chatify/middleware"
	authservice "github.com/abdullahshafaqat/Chatify/api/auth_service"
	messageservice "github.com/abdullahshafaqat/Chatify/api/message_service"
	"github.com/abdullahshafaqat/Chatify/db/firebase"
	"github.com/abdullahshafaqat/Chatify/db/mongodb"
	"github.com/abdullahshafaqat/Chatify/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	middleware.InitAuthMiddleware() // Initialize JWT secret

	firebase.InitFirebase()
	db := firebase.NewDB()
	authService := authservice.NewAuthService(db)

	mongodb.InitMongoDB()
	messageService := messageservice.NewMessageService()

	routerLayer := router.NewRouter(authService, messageService)

	r := gin.Default()
	routerLayer.DefineRoutes(r)

	log.Println("🚀 API Server is running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Failed to start API server: %v", err)
	}
}