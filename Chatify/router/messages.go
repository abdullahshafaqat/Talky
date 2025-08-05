package router

import (
	"net/http"
	"time"

	"github.com/abdullahshafaqat/Chatify/models"
	"github.com/gin-gonic/gin"
)
func (r *routerImpl) SendMessageHandler(c *gin.Context) {
	var input struct {
		ReceiverID string `json:"receiver_id" binding:"required"`
		Content    string `json:"content" binding:"required"`
	}

	// Bind only the receiver_id and content
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Get sender_id from context (set by AuthMiddleware)
	senderIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: user_id missing"})
		return
	}
	senderID := senderIDValue.(string)

	msg := &models.Message{
		SenderID:   senderID,
		ReceiverID: input.ReceiverID,
		Content:    input.Content,
		Timestamp:  time.Now(),
	}

	if err := r.messageService.SendMessage(msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "message sent"})
}
