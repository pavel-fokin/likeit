package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLikesGet(t *testing.T) {
	// setup
	req, _ := http.NewRequest("", "", nil)
	w := httptest.NewRecorder()

	// test
	LikesGet()(w, req)

	// assert
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
}

func TestLikesPost(t *testing.T) {
	// setup
	req, _ := http.NewRequest("", "", nil)
	w := httptest.NewRecorder()

	// test
	LikesPost()(w, req)

	// assert
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
}
