package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteMessageByID(messageID, senderID string) error {
	ctx := context.Background()

	objID, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return errors.New("invalid message ID")
	}

	filter := bson.M{"_id": objID, "sender_id": senderID}

	result, err := MessageCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("message not found or unauthorized")
	}
	return nil
}
