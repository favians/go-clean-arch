package usecase

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type jwtUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewJwtUsecase(u domain.UserRepository, to time.Duration) domain.JwtUsecase {
	return &jwtUsecase{
		userRepo:       u,
		contextTimeout: to,
	}
}

func (h *jwtUsecase) getOneUser(c context.Context, id string) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(c, h.contextTimeout)
	defer cancel()

	res, err := h.userRepo.GetOne(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}
