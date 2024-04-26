package app

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountLikes(t *testing.T) {
	// Setup.
	mockLikesDB := &mockLikeItDB{}
	likeIt := &App{
		db: mockLikesDB,
	}

	// Mock the behavior of the likesdb.Count method.
	mockLikesDB.On("CountLikes", context.Background()).Return(Likes(5), nil)

	// Test.
	count, err := likeIt.CountLikes(context.Background())

	// Assert the expected result.
	assert.Equal(t, Likes(5), count)
	assert.NoError(t, err)

	// Verify that the Count method was called.
	mockLikesDB.AssertCalled(t, "CountLikes", context.Background())
}

func TestIncrementLikes(t *testing.T) {
	// Setup.
	mockLikesDB := &mockLikeItDB{}
	likeIt := &App{
		db: mockLikesDB,
	}

	// Mock the behavior of the likesdb.Increment method.
	mockLikesDB.On("IncrementLikes", context.Background()).Return(nil)

	// Test.
	err := likeIt.IncrementLikes(context.Background())

	// Assert the expected result.
	assert.NoError(t, err)

	// Verify that the Increment method was called.
	mockLikesDB.AssertCalled(t, "IncrementLikes", context.Background())
}
