package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func (r *routerImpl) GetChatHandler(c *gin.Context) {
	userID := c.GetString("user_id") // get from JWT token

	messages, err := r.messageService.GetUserMessages(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

