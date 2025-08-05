package router

import (
	"net/http"

	"github.com/abdullahshafaqat/Chatify/models"
	"github.com/gin-gonic/gin"
)

func (r *routerImpl) UpdateMessageHandler(c *gin.Context) {
	// Extract user_id from token (set by middleware)
	senderID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req models.UpdateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := r.messageService.UpdateMessage(req.MessageID, senderID.(string), req.NewContent)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message updated successfully"})
}
