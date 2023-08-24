package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Likes struct {
	mock.Mock
}

func (m *Likes) Count(ctx context.Context) (int, error) {
	m.Called()
	return 0, nil
}

func (m *Likes) Increment(ctx context.Context) error {
	m.Called()
	return nil
}

func TestLikesGet(t *testing.T) {
	// Setup.
	req, _ := http.NewRequest("", "", nil)
	w := httptest.NewRecorder()

	likes := &Likes{}
	likes.On("Count").Return()

	// Test.
	LikesGet(likes)(w, req)

	// Assert.
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)

	likes.AssertNumberOfCalls(t, "Count", 1)
}

func TestLikesPost(t *testing.T) {
	// setup
	req, _ := http.NewRequest("", "", nil)
	w := httptest.NewRecorder()

	likes := &Likes{}
	likes.On("Increment").Return()

	// test
	LikesPost(likes)(w, req)

	// assert
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)

	likes.AssertNumberOfCalls(t, "Increment", 1)
}
