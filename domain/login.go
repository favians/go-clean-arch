package domain

import "context"

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUsecase interface {
	GetUser(ctx context.Context, username string, password string) (*User, error)
}
