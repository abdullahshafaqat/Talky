package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateMessageByID(messageID, senderID, newContent string) error {
	ctx := context.Background()

	objID, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return errors.New("invalid message ID")
	}

	filter := bson.M{"_id": objID, "sender_id": senderID}
	update := bson.M{
		"$set": bson.M{
			"content":    newContent,
			"updated_at": time.Now(),
		},
	}

	result, err := MessageCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("message not found or unauthorized")
	}
	return nil
}
