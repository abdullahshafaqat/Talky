package mongodb

import (
	"context"
	"time"

	"github.com/abdullahshafaqat/Chatify/models"
)

func SaveMessage(msg *models.Message) error {
	msg.Timestamp = time.Now()

	_, err := MessageCollection.InsertOne(context.Background(), msg)
	return err
}
