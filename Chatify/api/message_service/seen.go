package messageservice

import (
	"github.com/abdullahshafaqat/Chatify/db/mongodb"
)

func (s *serviceImpl) MarkMessageSeen(messageID, userID string) error {
	return mongodb.UpdateMessageSeenStatus(messageID, userID)
}

func (s *serviceImpl) MarkMessageDelivered(messageID, userID string) error {
	return mongodb.UpdateMessageDeliveredStatus(messageID, userID)
}
