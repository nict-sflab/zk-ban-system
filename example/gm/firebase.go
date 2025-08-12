package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/akakou/zk-ban-system/core"
	"google.golang.org/api/option"
)

type FirebaseAuthService struct {
	authClient *auth.Client
}

func NewFirebaseAuthService(ctx context.Context, credentialsFile string) (*FirebaseAuthService, error) {
	var opts []option.ClientOption
	if credentialsFile != "" {
		opts = append(opts, option.WithCredentialsFile(credentialsFile))
	}

	app, err := firebase.NewApp(ctx, nil, opts...)
	if err != nil {
		return nil, fmt.Errorf("firebase app init error: %w", err)
	}

	ac, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("firebase auth client init error: %w", err)
	}

	return &FirebaseAuthService{authClient: ac}, nil
}

func (s *FirebaseAuthService) FirebaseAuth() func(t *core.JoinRequest[string]) (string, error) {
	return func(t *core.JoinRequest[string]) (string, error) {
		if s.authClient == nil {
			return "", fmt.Errorf("FirebaseAuthService not initialized")
		}

		ctx := context.Background()
		idToken := t.Option
		if idToken == "" {
			return "", fmt.Errorf("missing ID token")
		}

		token, err := s.authClient.VerifyIDToken(ctx, idToken)
		if err != nil {
			return "", fmt.Errorf("invalid ID token: %w", err)
		}

		if v, ok := token.Claims["phone_number"].(string); ok && v != "" {
			return v, nil
		}

		user, err := s.authClient.GetUser(ctx, token.UID)
		if err != nil {
			return "", fmt.Errorf("failed to get user: %w", err)
		}
		if user.PhoneNumber == "" {
			return "", fmt.Errorf("no phone number found")
		}
		return user.PhoneNumber, nil
	}
}
