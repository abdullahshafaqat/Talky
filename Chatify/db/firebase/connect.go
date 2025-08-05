package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/abdullahshafaqat/Chatify/config"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() {
	// Get Firebase key path from config
	credPath := config.GetFirebaseKeyPath()

	opt := option.WithCredentialsFile(credPath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("❌ Firebase initialization error: %v", err)
	}

	App = app
}
