package authservice

import (
	"github.com/abdullahshafaqat/Chatify/db/firebase"
	"github.com/gin-gonic/gin"
)

func (s *serviceImpl) VerifyOTP(c *gin.Context, email, otp string) (string, error) {
	return firebase.VerifyOTP(email, otp)
}
