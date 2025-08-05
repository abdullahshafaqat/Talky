package firebase

import (
	"context"
	"errors"
	"time"

	"github.com/abdullahshafaqat/Chatify/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func (d *dbImpl) CreateUser(c *gin.Context, user *models.UserSignup) error {
	ctx := context.Background()

	client, err := App.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// Check if user exists
	iter := client.Collection("users").Where("email", "==", user.Email).Documents(ctx)
	if _, err := iter.Next(); err != iterator.Done {
		return errors.New("user already exists with this email")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	doc, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"username":     user.Username,
		"email":        user.Email,
		"photo_url":    user.PhotoURL,
		"phone_number": user.PhoneNumber,
		"created_at":   user.CreatedAt,
		"updated_at":   user.UpdatedAt,
	})

	if err != nil {
		return err
	}

	user.ID = doc.ID
	return nil
}
