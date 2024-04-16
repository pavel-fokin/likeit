package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

// ErrorResponse.
type ErrorResponse struct {
	Data struct {
		Errors []Error `json:"errors"`
	} `json:"data"`
}

// AsErrorResponse.
func asErrorResponse(
	w http.ResponseWriter, err error, statusCode int,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Make error response.
	payload := ErrorResponse{}
	payload.Data.Errors = []Error{{Message: fmt.Sprint(err)}}

	// Encode json.
	json.NewEncoder(w).Encode(payload)
}

// AsSuccessResponse.
func asSuccessResponse(
	w http.ResponseWriter, payload interface{}, statusCode int,
) {
	if payload != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	}

	w.WriteHeader(statusCode)
}
