package messageservice

import "github.com/abdullahshafaqat/Chatify/db/mongodb"

func (s *serviceImpl) UpdateMessage(messageID, senderID, newContent string) error {
	return mongodb.UpdateMessageByID(messageID, senderID, newContent)
}
