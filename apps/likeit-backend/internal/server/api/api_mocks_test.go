package api

import (
	"context"
	"pavel-fokin/likeit/internal/app"

	"github.com/stretchr/testify/mock"
)

type AuthMock struct {
	mock.Mock
}

func (m *AuthMock) SignIn(ctx context.Context, username, password string) (*app.User, error) {
	args := m.Called(ctx, username, password)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*app.User), args.Error(1)
}

func (m *AuthMock) SignUp(ctx context.Context, username, password string) (*app.User, error) {
	args := m.Called(ctx, username, password)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*app.User), args.Error(1)
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
