package auth

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func newClient(ctx context.Context) (*auth.Client, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	return app.Auth(ctx)
}

func Authenticate(ctx context.Context, token string) (uid string, err error) {
	client, err := newClient(ctx)
	if err != nil {
		return "", err
	}
	fbToken, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		return "", err
	}
	uid = fbToken.UID
	return
}
