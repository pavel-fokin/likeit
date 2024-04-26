package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"pavel-fokin/likeit/internal/app"
	"pavel-fokin/likeit/internal/server/apiutil"
)

type Auth interface {
	SignIn(ctx context.Context, username, password string) (*app.User, error)
	SignUp(ctx context.Context, username, password string) (*app.User, error)
}

// SignIn signs in a user.
func SignIn(app Auth, tokenSigningKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req SignInRequest
		if err := apiutil.ParseJSON(r, &req); err != nil {
			slog.ErrorContext(ctx, "failed to parse request body", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		user, err := app.SignIn(ctx, req.Username, req.Password)
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

		apiutil.AsSuccessResponse(w, &SignInResponse{
			AccessToken: accessToken,
		}, http.StatusNoContent)
	}
}

// SignUp signs up a user.
func SignUp(app Auth, tokenSigningKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req SignUpRequest
		if err := apiutil.ParseJSON(r, &req); err != nil {
			slog.ErrorContext(ctx, "failed to parse request body", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		user, err := app.SignUp(ctx, req.Username, req.Password)
		if err != nil {
			slog.ErrorContext(ctx, "failed to sign up user", "err", err)
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
