package models

import "time"

type UserSignup struct {
	ID        string    `firestore:"id,omitempty"`
	Username  string    `json:"username" firestore:"username"`
	Email     string    `json:"email" firestore:"email"`
	PhotoURL  string    `json:"photo_url" firestore:"photo_url"`
	PhoneNumber string    `json:"phone_number" firestore:"phone_number"`
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}
