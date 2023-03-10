package service

import (
	"context"

	"github.com/learningPlatform/internal/models"
)

type User struct {
	userStorage UserStorage
}

func NewUser(userStorage UserStorage) *User {
	return &User{
		userStorage: userStorage,
	}
}

func (u User) CreateUser(ctx context.Context, user models.User) (int, error) {
	id, err := u.userStorage.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
