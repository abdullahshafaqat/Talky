// router/login_handler.go
package router

import (
	"net/http"

	"github.com/abdullahshafaqat/Chatify/models"
	"github.com/gin-gonic/gin"
)

func (r *routerImpl) LoginHandler(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}

	otp, err := r.authService.Login(c, req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP sent to email",
		"otp":     otp, // ✅ include OTP in response
	})
}
