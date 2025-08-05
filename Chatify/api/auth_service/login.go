package authservice

import (
	"github.com/abdullahshafaqat/Chatify/db/firebase"
	"github.com/gin-gonic/gin"
)

func (s *serviceImpl) Login(c *gin.Context, email string) (string, error) {
	return firebase.SendLoginOTP(email)
}
