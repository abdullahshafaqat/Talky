package firebase

import (
	"github.com/abdullahshafaqat/Chatify/models"
	"github.com/gin-gonic/gin"
)

type DB interface {
	CreateUser(c *gin.Context, user *models.UserSignup) error
	GetAllUsers(c *gin.Context) ([]*models.UserSignup, error) // Add this line
}

type dbImpl struct {}

func NewDB() DB {
	return &dbImpl{}
}