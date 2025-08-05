package authservice

import (
	"github.com/abdullahshafaqat/Chatify/models"
	"github.com/gin-gonic/gin"
)

func (s *serviceImpl) GetAllUsers(c *gin.Context) ([]*models.UserSignup, error) {
	return s.db.GetAllUsers(c)
}