package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLikes(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		// Setup.
		req, _ := http.NewRequest("", "", nil)
		w := httptest.NewRecorder()

		likes := &LikesMock{}
		likes.On("CountLikes").Return(nil)

		// Test.
		GetLikes(likes)(w, req)

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
		GetLikes(likes)(w, req)

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
		PostLikes(likes)(w, req)

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
		PostLikes(likes)(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(500, resp.StatusCode)

		likes.AssertNumberOfCalls(t, "IncrementLikes", 1)
	})
}
