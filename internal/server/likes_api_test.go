package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pavel-fokin/likeit/internal/app"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetLikes(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		// Setup.
		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		likes := &LikesMock{}
		likes.On("CountLikes").Return(nil)

		// Test.
		getLikes(likes)(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		likes.AssertNumberOfCalls(t, "CountLikes", 1)
	})

	t.Run("Failure", func(t *testing.T) {
		// Setup.
		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		likes := &LikesMock{}
		likes.On("CountLikes").Return(fmt.Errorf("some error"))

		// Test.
		getLikes(likes)(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		likes.AssertNumberOfCalls(t, "CountLikes", 1)
	})
}

func TestLikesPost(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		assert := assert.New(t)

		// Setup.
		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		likes := &LikesMock{}
		likes.On("IncrementLikes").Return(nil)

		// Test.
		postLikes(likes)(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(204, resp.StatusCode)

		likes.AssertNumberOfCalls(t, "IncrementLikes", 1)
	})

	t.Run("Failure", func(t *testing.T) {
		assert := assert.New(t)

		// Setup.
		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		likes := &LikesMock{}
		likes.On("IncrementLikes").Return(fmt.Errorf("error"))

		// Test.
		postLikes(likes)(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(500, resp.StatusCode)

		likes.AssertNumberOfCalls(t, "IncrementLikes", 1)
	})
}

type LikesMock struct {
	mock.Mock
}

func (m *LikesMock) CountLikes(ctx context.Context) (app.Likes, error) {
	args := m.Called()
	return app.Likes(0), args.Error(0)
}

func (m *LikesMock) IncrementLikes(ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}
