package firebase

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"google.golang.org/api/iterator"
)

func GenerateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000)) // 6-digit OTP
}

func SendOTPToEmail(email, otp string) {
	// You can integrate with real SMTP service here
	fmt.Printf("📧 Sending OTP %s to email: %s\n", otp, email)
}

func SendLoginOTP(email string) (string, error) {
	ctx := context.Background()

	client, err := App.Firestore(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	iter := client.Collection("users").Where("email", "==", email).Documents(ctx)
	_, err = iter.Next()
	if err == iterator.Done {
		return "", errors.New("user not found with this email")
	}
	if err != nil {
		return "", err
	}

	otp := GenerateOTP()
	SaveOTP(email, otp)
	SendOTPToEmail(email, otp)

	return otp, nil
}

func VerifyOTP(email, inputOtp string) (string, error) {
	expectedOtp := GetOTP(email)
	if expectedOtp == "" {
		return "", errors.New("OTP expired or not found")
	}
	if inputOtp != expectedOtp {
		return "", errors.New("invalid OTP")
	}
	DeleteOTP(email)

	ctx := context.Background()
	client, err := App.Firestore(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	iter := client.Collection("users").Where("email", "==", email).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done || err != nil {
		return "", errors.New("user not found")
	}

	return doc.Ref.ID, nil
}
