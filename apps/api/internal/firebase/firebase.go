package firebaseapp

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	fbauth "firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func InitializeAuthClient(ctx context.Context) (*fbauth.Client, error) {
	credentialsFile := credentialsPath()
	if credentialsFile == "" {
		return nil, fmt.Errorf("firebase credentials path is not configured")
	}

	opt := option.WithCredentialsFile(credentialsFile)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("initialize app: %w", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("initialize auth client: %w", err)
	}

	log.Println("firebase connection established")

	return authClient, nil
}

func credentialsPath() string {
	if value := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); value != "" {
		return value
	}

	if value := os.Getenv("FIREBASE_SERVICE_ACCOUNT_JSON"); value != "" {
		return value
	}

	return ""
}