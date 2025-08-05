package authservice

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/abdullahshafaqat/Chatify/models"
)

func (s *serviceImpl) SignUp(c *gin.Context, user *models.UserSignup) error {
	if !isValidEmail(user.Email) {
		return errors.New("please enter a valid Gmail address")
	}
	return s.db.CreateUser(c, user)
}

func isValidEmail(email string) bool {
	return regexp.MustCompile(`^[^@]+@gmail\.com$`).MatchString(email)
}
