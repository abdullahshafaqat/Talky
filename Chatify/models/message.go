package models

import "time"

type Message struct {
	ID         string    `bson:"_id,omitempty" json:"id"`
	SenderID   string    `bson:"sender_id" json:"sender_id"`
	ReceiverID string    `bson:"receiver_id" json:"receiver_id"`
	Content    string    `bson:"content" json:"content"`
	Seen       bool      `bson:"seen" json:"seen"`
	Delivered  bool      `bson:"delivered" json:"delivered"`
	Timestamp  time.Time `bson:"timestamp" json:"timestamp"`
}

type UpdateMessageRequest struct {
	MessageID  string `json:"message_id"`
	NewContent string `json:"new_content"`
}

type DeleteMessageRequest struct {
	MessageID string `json:"message_id"`
}