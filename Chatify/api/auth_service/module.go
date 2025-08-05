package authservice

import (
	"github.com/abdullahshafaqat/Chatify/db/firebase"
	"github.com/abdullahshafaqat/Chatify/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	SignUp(c *gin.Context, user *models.UserSignup) error
	Login(c *gin.Context, email string) (string, error)
	VerifyOTP(c *gin.Context, email, otp string) (string, error)
	GetAllUsers(c *gin.Context) ([]*models.UserSignup, error) // Add this line
}

type serviceImpl struct {
	db firebase.DB
}

func NewAuthService(db firebase.DB) Service {
	return &serviceImpl{db: db}
}