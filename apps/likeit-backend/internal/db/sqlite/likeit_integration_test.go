package sqlite

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"pavel-fokin/likeit/internal/app"
)

func TestLikeItDB(t *testing.T) {
	likeitDB, close := New(":memory:")
	defer close()

	err := likeitDB.IncrementLikes(context.Background())
	assert.NoError(t, err)

	count, err := likeitDB.CountLikes(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, count, app.Likes(1))
}

func TestCreateUser(t *testing.T) {
	likeitDB, close := New(":memory:")
	defer close()

	user, err := likeitDB.CreateUser(context.Background(), "username", "password")
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestFindUser(t *testing.T) {
	likeitDB, close := New(":memory:")
	defer close()

	user, err := likeitDB.CreateUser(context.Background(), "username", "password")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	foundUser, err := likeitDB.FindUser(context.Background(), "username")
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.ID, foundUser.ID)
}
