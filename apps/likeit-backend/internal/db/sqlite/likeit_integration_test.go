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
