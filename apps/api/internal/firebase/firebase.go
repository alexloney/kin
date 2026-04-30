package firebaseapp

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func Initialize(ctx context.Context) error {
	credentialsFile := credentialsPath()
	if credentialsFile == "" {
		return fmt.Errorf("firebase credentials path is not configured")
	}

	opt := option.WithCredentialsFile(credentialsFile)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return fmt.Errorf("initialize app: %w", err)
	}

	if _, err := app.Auth(ctx); err != nil {
		return fmt.Errorf("initialize auth client: %w", err)
	}

	log.Println("firebase connection established")

	return nil
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