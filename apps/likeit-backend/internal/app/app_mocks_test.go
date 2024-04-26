package app

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// Mock implementation of the DB interface.
type mockLikeItDB struct {
	mock.Mock
}

var _ DB = (*mockLikeItDB)(nil)

func (m *mockLikeItDB) CountLikes(ctx context.Context) (Likes, error) {
	args := m.Called(ctx)
	return args.Get(0).(Likes), args.Error(1)
}

func (m *mockLikeItDB) IncrementLikes(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *mockLikeItDB) CreateUser(ctx context.Context, username, password string) (*User, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockLikeItDB) FindUser(ctx context.Context, username string) (*User, error) {
	args := m.Called(ctx, username)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*User), args.Error(1)
}
