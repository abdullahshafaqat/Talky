package messageservice

import (
	"github.com/abdullahshafaqat/Chatify/models"
)

type Service interface {
	SendMessage(msg *models.Message) error
	GetUserMessages(userID string) ([]models.Message, error)
	UpdateMessage(messageID, senderID, newContent string) error
	DeleteMessage(messageID, senderID string) error
	MarkMessageSeen(messageID, userID string) error
	MarkMessageDelivered(messageID, userID string) error
}

type serviceImpl struct{}

func NewMessageService() Service {
	return &serviceImpl{}
}
