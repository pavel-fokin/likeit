package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"pavel-fokin/likeit/internal/app"
	"pavel-fokin/likeit/internal/server/apiutil"
)

type AppAuth interface {
	SignIn(ctx context.Context, userID string) (*app.User, error)
	SignUp(ctx context.Context) (*app.User, error)
}

// SignIn signs in a user.
func SignIn(app AppAuth, tokenSigningKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req SignInRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			slog.ErrorContext(ctx, "failed to decode request body", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		user, err := app.SignIn(ctx, req.UserID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to sign in user", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		accessToken, err := apiutil.NewAccessToken(user.ID, tokenSigningKey)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create access token", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, SignInResponse{
			AccessToken: accessToken,
		}, http.StatusNoContent)
	}
}

// SignUp signs up a user.
func SignUp(app AppAuth, tokenSigningKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user, err := app.SignUp(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to sign up user: %v", err), http.StatusInternalServerError)
			return
		}

		accessToken, err := apiutil.NewAccessToken(user.ID, tokenSigningKey)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create access token", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, SignInResponse{
			AccessToken: accessToken,
		}, http.StatusNoContent)
	}
}
