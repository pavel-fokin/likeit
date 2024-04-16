package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type LikesMock struct {
	mock.Mock
}

func (m *LikesMock) CountLikes(ctx context.Context) (int, error) {
	args := m.Called()
	return 0, args.Error(0)
}

func (m *LikesMock) IncrementLikes(ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}

func TestLikesGet_Success(t *testing.T) {
	// Setup.
	req, _ := http.NewRequest("", "", nil)
	w := httptest.NewRecorder()

	likes := &LikesMock{}
	likes.On("CountLikes").Return(nil)

	// Test.
	LikesGet(likes)(w, req)

	// Assert.
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)

	likes.AssertNumberOfCalls(t, "CountLikes", 1)
}

func TestLikesGet_Failure(t *testing.T) {
	// Setup.
	req, _ := http.NewRequest("", "", nil)
	w := httptest.NewRecorder()

	likes := &LikesMock{}
	likes.On("CountLikes").Return(fmt.Errorf("error"))

	// Test.
	LikesGet(likes)(w, req)

	// Assert.
	resp := w.Result()
	assert.Equal(t, 500, resp.StatusCode)

	likes.AssertNumberOfCalls(t, "CountLikes", 1)
}

func TestLikesPost_Success(t *testing.T) {
	assert := assert.New(t)

	// Setup.
	req, _ := http.NewRequest("", "", nil)
	w := httptest.NewRecorder()

	likes := &LikesMock{}
	likes.On("IncrementLikes").Return(nil)

	// Test.
	LikesPost(likes)(w, req)

	// Assert.
	resp := w.Result()
	assert.Equal(204, resp.StatusCode)

	likes.AssertNumberOfCalls(t, "IncrementLikes", 1)
}

func TestLikesPost_Failure(t *testing.T) {
	assert := assert.New(t)

	// Setup.
	req, _ := http.NewRequest("", "", nil)
	w := httptest.NewRecorder()

	likes := &LikesMock{}
	likes.On("IncrementLikes").Return(fmt.Errorf("error"))

	// Test.
	LikesPost(likes)(w, req)

	// Assert.
	resp := w.Result()
	assert.Equal(500, resp.StatusCode)

	likes.AssertNumberOfCalls(t, "IncrementLikes", 1)
}
