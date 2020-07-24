package usecase

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/bootstrap"
	"github.com/bxcodec/go-clean-arch/domain"
	_userRepo "github.com/bxcodec/go-clean-arch/user/repository/mongo"
)

type jwtUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewJwtUsecase() domain.JwtUsecase {
	return &jwtUsecase{
		userRepo:       _userRepo.NewMongoRepository(bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongo.name"))),
		contextTimeout: time.Duration(bootstrap.App.Config.GetInt("context.timeout")) * time.Second,
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
