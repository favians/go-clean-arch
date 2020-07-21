package usecase

import (
	"context"
	"log"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type loginUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(u domain.UserRepository, to time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepo:       u,
		contextTimeout: to,
	}
}

func (login *loginUsecase) GetUser(c context.Context, username string, password string) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(c, login.contextTimeout)
	defer cancel()

	res, err := login.userRepo.GetByCredential(ctx, username, password)
	if err != nil {
		return res, err
	}
	log.Println(res)

	return res, nil
}
