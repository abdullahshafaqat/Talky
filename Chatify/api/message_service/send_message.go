package messageservice

import (
	"github.com/abdullahshafaqat/Chatify/db/mongodb"
	"github.com/abdullahshafaqat/Chatify/models"
)

func (s *serviceImpl) SendMessage(msg *models.Message) error {
	return mongodb.SaveMessage(msg)
}
