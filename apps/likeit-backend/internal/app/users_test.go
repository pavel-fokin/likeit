package app

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	db := &mockLikeItDB{}
	db.On("CreateUser", context.Background()).Return(&User{}, nil)

	app := &App{
		db: db,
	}

	user, err := app.SignUp(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestSignIn(t *testing.T) {
	t.Run("user not found", func(t *testing.T) {
		db := &mockLikeItDB{}
		db.On("FindUser", context.Background(), "userID").Return(nil, errors.New("user not found"))

		app := &App{
			db: db,
		}

		user, err := app.SignIn(context.Background(), "userID")
		assert.ErrorContains(t, err, "failed to sign in user: user not foun")
		assert.Nil(t, user)
	})

	t.Run("user found", func(t *testing.T) {
		db := &mockLikeItDB{}
		db.On("FindUser", context.Background(), "userID").Return(&User{}, nil)

		app := &App{
			db: db,
		}

		user, err := app.SignIn(context.Background(), "userID")
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})
}
