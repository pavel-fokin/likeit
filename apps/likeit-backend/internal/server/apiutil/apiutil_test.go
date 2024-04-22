package apiutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestAsSuccessResponse(t *testing.T) {
	// Create a mock HTTP response writer.
	w := httptest.NewRecorder()

	// Define the payload and status code.
	payload := map[string]interface{}{
		"message": "Success",
	}
	statusCode := http.StatusOK

	// Call the function being tested.
	AsSuccessResponse(w, payload, statusCode)

	// Verify the response.
	response := w.Result()
	defer response.Body.Close()

	assert.Equal(t, statusCode, response.StatusCode)

	// Verify the payload
	var responseBody map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	assert.NoError(t, err)
	assert.Equal(t, payload, responseBody)
}
