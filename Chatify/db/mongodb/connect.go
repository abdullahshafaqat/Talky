package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/abdullahshafaqat/Chatify/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MessageCollection *mongo.Collection

func InitMongoDB() {
	// Get MongoDB URI from config
	mongoURI := config.GetMongoURI()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("❌ MongoDB connection failed: %v", err)
	}

	MongoClient = client
	MessageCollection = client.Database("chatify").Collection("messages")
	log.Println("✅ MongoDB connected")
}
