package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteMessageRequest struct {
	MessageID string `json:"message_id"`
}

func (r *routerImpl) DeleteMessageHandler(c *gin.Context) {
	senderID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req DeleteMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.MessageID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := r.messageService.DeleteMessage(req.MessageID, senderID.(string))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}
