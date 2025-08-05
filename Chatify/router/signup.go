package router

import (
	"net/http"

	"github.com/abdullahshafaqat/Chatify/models"
	"github.com/gin-gonic/gin"
)

func (r *routerImpl) SignUpHandler(c *gin.Context) {
	var newUser models.UserSignup

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := r.authService.SignUp(c, &newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"info": gin.H{
			"id":           newUser.ID,
			"username":     newUser.Username,
			"email":        newUser.Email,
			"photo_url":    newUser.PhotoURL,
			"created_at":   newUser.CreatedAt,
			"updated_at":   newUser.UpdatedAt,
			"phone_number": newUser.PhoneNumber,
		},
	})
}
