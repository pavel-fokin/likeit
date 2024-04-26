package app

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pavel-fokin/likeit/internal/app/apputil"
)

func TestSignUp(t *testing.T) {
	db := &mockLikeItDB{}
	db.On("CreateUser", context.Background(), "username", mock.Anything).Return(&User{}, nil)

	app := &App{
		db: db,
	}

	user, err := app.SignUp(context.Background(), "username", "password")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	// Verify that password is not stored in plain text.
	assert.NotEqual(t, "password", db.Calls[0].Arguments[2])
}

func TestSignIn(t *testing.T) {
	t.Run("user not found", func(t *testing.T) {
		db := &mockLikeItDB{}
		db.On("FindUser", context.Background(), mock.Anything).Return(nil, errors.New("user not found"))

		app := &App{
			db: db,
		}

		user, err := app.SignIn(context.Background(), "username", "password")
		assert.ErrorContains(t, err, "failed to sign in user: user not found")
		assert.Nil(t, user)
	})

	t.Run("user found", func(t *testing.T) {
		hashedPassword, err := apputil.HashPassword("password")
		assert.NoError(t, err)

		db := &mockLikeItDB{}
		db.On("FindUser", context.Background(), mock.Anything).Return(&User{Password: hashedPassword}, nil)

		app := &App{
			db: db,
		}

		user, err := app.SignIn(context.Background(), "username", "password")
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})
}
