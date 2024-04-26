package apiutil

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidation() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

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
func AsErrorResponse(
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
func AsSuccessResponse(
	w http.ResponseWriter, payload any, statusCode int,
) {
	if payload != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	}

	w.WriteHeader(statusCode)
}

// ParseJSON parses the request body into the given interface.
func ParseJSON(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	if err := validate.Struct(v); err != nil {
		return err
	}

	return nil
}
