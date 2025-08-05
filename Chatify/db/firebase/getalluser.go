package firebase

import (
	"context"
	"time"

	"github.com/abdullahshafaqat/Chatify/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func (d *dbImpl) GetAllUsers(c *gin.Context) ([]*models.UserSignup, error) {
	ctx := context.Background()
	client, err := App.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	var users []*models.UserSignup
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		user := &models.UserSignup{
			ID:          doc.Ref.ID,
			Username:    doc.Data()["username"].(string),
			Email:       doc.Data()["email"].(string),
			PhotoURL:    doc.Data()["photo_url"].(string),
			PhoneNumber: doc.Data()["phone_number"].(string),
			CreatedAt:   doc.Data()["created_at"].(time.Time),
			UpdatedAt:   doc.Data()["updated_at"].(time.Time),
		}
		users = append(users, user)
	}

	return users, nil
}
