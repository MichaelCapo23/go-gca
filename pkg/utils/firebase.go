package utils

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
)

// InitializeFirebase initializes the firebase
func InitializeFirebase(json string) (*firebase.App, error) {
	opt := option.WithCredentialsJSON([]byte(json))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
