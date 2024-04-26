package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"pavel-fokin/likeit/internal/app"
)

func TestSignIn(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Setup.
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/signin", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On("SignIn", context.Background(), "username", "password").Return(&app.User{}, nil)

		// Test.
		SignIn(auth, "token")(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "SignIn", 1)
	})

	t.Run("Failure", func(t *testing.T) {
		// Setup.
		body := `{"username": "username", "password": "password"}`
		req, _ := http.NewRequest("POST", "/signup", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}
		auth.On("SignIn", context.Background(), "username", "password").Return(nil, fmt.Errorf("some error"))

		// Test.
		SignIn(auth, "token")(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(t, 500, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "SignIn", 1)
	})

	t.Run("Invalid request", func(t *testing.T) {
		// Setup.
		body := `{"username": "username"}`
		req, _ := http.NewRequest("POST", "/signin", strings.NewReader(body))
		w := httptest.NewRecorder()

		auth := &AuthMock{}

		// Test.
		SignIn(auth, "token")(w, req)

		// Assert.
		resp := w.Result()
		assert.Equal(t, 400, resp.StatusCode)

		auth.AssertNumberOfCalls(t, "SignIn", 0)
	})
}
