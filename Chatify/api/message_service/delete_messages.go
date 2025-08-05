package messageservice

import "github.com/abdullahshafaqat/Chatify/db/mongodb"

func (s *serviceImpl) DeleteMessage(messageID, senderID string) error {
	return mongodb.DeleteMessageByID(messageID, senderID)
}
