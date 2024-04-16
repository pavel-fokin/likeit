package likeit

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCountLikes(t *testing.T) {
	// Setup.
	mockLikesDB := &mockLikesDB{}
	likeIt := &LikeIt{
		likesdb: mockLikesDB,
	}

	// Mock the behavior of the likesdb.Count method.
	mockLikesDB.On("Count", context.Background()).Return(5, nil)

	// Test.
	count, err := likeIt.CountLikes(context.Background())

	// Assert the expected result.
	assert.Equal(t, 5, count)
	assert.NoError(t, err)

	// Verify that the Count method was called.
	mockLikesDB.AssertCalled(t, "Count", context.Background())
}

func TestIncrementLikes(t *testing.T) {
	// Setup.
	mockLikesDB := &mockLikesDB{}
	likeIt := &LikeIt{
		likesdb: mockLikesDB,
	}

	// Mock the behavior of the likesdb.Increment method.
	mockLikesDB.On("Increment", context.Background()).Return(nil)

	// Test.
	err := likeIt.IncrementLikes(context.Background())

	// Assert the expected result.
	assert.NoError(t, err)

	// Verify that the Increment method was called.
	mockLikesDB.AssertCalled(t, "Increment", context.Background())
}

// Mock implementation of the LikesDB interface.
type mockLikesDB struct {
	mock.Mock
}

func (m *mockLikesDB) Count(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

func (m *mockLikesDB) Increment(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
