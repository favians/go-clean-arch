package usecase

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/spf13/viper"
)

type jwtUsecase struct {
	UserRepo       domain.UserRepository
	ContextTimeout time.Duration
	Config         *viper.Viper
}

func NewJwtUsecase(u domain.UserRepository, to time.Duration, config *viper.Viper) domain.JwtUsecase {
	return &jwtUsecase{
		UserRepo:       u,
		ContextTimeout: to,
		Config:         config,
	}
}

func (h *jwtUsecase) getOneUser(c context.Context, id string) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(c, h.ContextTimeout)
	defer cancel()

	res, err := h.UserRepo.GetOne(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}
