package messageservice

import (
	"github.com/abdullahshafaqat/Chatify/db/mongodb"
	"github.com/abdullahshafaqat/Chatify/models"
)

func (s *serviceImpl) GetUserMessages(userID string) ([]models.Message, error) {
	return mongodb.GetMessagesForUser(userID)
}
