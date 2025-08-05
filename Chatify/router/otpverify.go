package router

import (
	"net/http"

	"github.com/abdullahshafaqat/Chatify/middleware"
	"github.com/abdullahshafaqat/Chatify/models"
	"github.com/gin-gonic/gin"
)

func (r *routerImpl) VerifyOTPHandler(c *gin.Context) {
	var req models.OTPVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userID, err := r.authService.VerifyOTP(c, req.Email, req.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// ✅ Generate JWT token with user ID
	token, err := middleware.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	// ✅ Return token in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user_id": userID,
		"token":   token,
	})
}
