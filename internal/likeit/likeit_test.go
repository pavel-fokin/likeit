package likeit

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCountLikes(t *testing.T) {
	// Setup.
	mockLikesDB := &mockLikeItDB{}
	likeIt := &LikeIt{
		db: mockLikesDB,
	}

	// Mock the behavior of the likesdb.Count method.
	mockLikesDB.On("CountLikes", context.Background()).Return(5, nil)

	// Test.
	count, err := likeIt.CountLikes(context.Background())

	// Assert the expected result.
	assert.Equal(t, 5, count)
	assert.NoError(t, err)

	// Verify that the Count method was called.
	mockLikesDB.AssertCalled(t, "CountLikes", context.Background())
}

func TestIncrementLikes(t *testing.T) {
	// Setup.
	mockLikesDB := &mockLikeItDB{}
	likeIt := &LikeIt{
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

// Mock implementation of the LikesDB interface.
type mockLikeItDB struct {
	mock.Mock
}

func (m *mockLikeItDB) CountLikes(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

func (m *mockLikeItDB) IncrementLikes(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
