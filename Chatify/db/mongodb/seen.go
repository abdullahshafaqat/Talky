package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateMessageSeenStatus(messageID, userID string) error {
	ctx := context.Background()
	objID, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID, "receiver_id": userID}
	update := bson.M{"$set": bson.M{"seen": true}}

	result, err := MessageCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no matching message found")
	}
	return nil
}

func UpdateMessageDeliveredStatus(messageID, userID string) error {
	ctx := context.Background()
	objID, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID, "receiver_id": userID}
	update := bson.M{"$set": bson.M{"delivered": true}}

	result, err := MessageCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no matching message found")
	}
	return nil
}
