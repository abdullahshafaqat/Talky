package mongodb

import (
	"context"

	"github.com/abdullahshafaqat/Chatify/models"
)
func GetMessagesForUser(userID string) ([]models.Message, error) {
	ctx := context.Background()

	filter := map[string]interface{}{
		"$or": []map[string]interface{}{
			{"sender_id": userID},
			{"receiver_id": userID},
		},
	}

	cursor, err := MessageCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []models.Message
	for cursor.Next(ctx) {
		var msg models.Message
		if err := cursor.Decode(&msg); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
