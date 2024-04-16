package sqlite

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"pavel-fokin/likeit/internal/app"
)

func TestLikesCount(t *testing.T) {
	likes, close := New(":memory:")
	defer close()

	// Test.
	count, err := likes.CountLikes(context.Background())

	assert.Equal(t, count, app.Likes(0))
	assert.NoError(t, err)
}

func TestLikesIncrement(t *testing.T) {
	likes, close := New(":memory:")
	defer close()

	// Test.
	err := likes.IncrementLikes(context.Background())

	assert.NoError(t, err)
}
