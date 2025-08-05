package router

import (
	"github.com/abdullahshafaqat/Chatify/middleware"
	"github.com/gin-gonic/gin"
)

func (r *routerImpl) DefineRoutes(router *gin.Engine) {
	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	api := router.Group("/api")
	{
		// Public routes
		api.POST("/signup", r.SignUpHandler)
		api.POST("/login", r.LoginHandler)
		api.POST("/verify-otp", r.VerifyOTPHandler)

		// Protected routes
		api.Use(middleware.AuthMiddleware())
		{
			api.POST("/send", r.SendMessageHandler)
			api.GET("/messages", r.GetChatHandler)
			api.PUT("/update-message", r.UpdateMessageHandler)
			api.DELETE("/delete-message", r.DeleteMessageHandler)
			api.GET("/users", r.GetUsersHandler) // Add this line
		}
	}
}