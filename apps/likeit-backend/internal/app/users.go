package app

import (
	"context"
	"fmt"
)

// User is an application model that represents a user.
type User struct {
	ID string
}

// SignIn signs in a user.
func (a *App) SignIn(ctx context.Context, username, password string) (*User, error) {
	user, err := a.db.FindUser(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to sign in user: %w", err)
	}
	return user, nil
}

// SignUp signs up a user.
func (a *App) SignUp(ctx context.Context, username, password string) (*User, error) {
	user, err := a.db.CreateUser(ctx, username, password)
	if err != nil {
		return nil, fmt.Errorf("failed to sign up user: %w", err)
	}
	return user, nil
}
