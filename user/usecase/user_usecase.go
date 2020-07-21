package usecase

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepository, to time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: to,
	}
}

func (user *userUsecase) Store(c context.Context, m *domain.User) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(c, user.contextTimeout)
	defer cancel()

	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	res, err := user.userRepo.Store(ctx, m)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (user *userUsecase) GetOne(c context.Context, id string) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(c, user.contextTimeout)
	defer cancel()

	res, err := user.userRepo.GetOne(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (user *userUsecase) GetAllWithPage(c context.Context, rp int64, p int64, filter interface{}, setsort interface{}) ([]domain.User, int64, error) {

	ctx, cancel := context.WithTimeout(c, user.contextTimeout)
	defer cancel()

	res, count, err := user.userRepo.GetAllWithPage(ctx, rp, p, filter, setsort)
	if err != nil {
		return res, count, err
	}

	return res, count, nil
}

func (user *userUsecase) Update(c context.Context, m *domain.User, id string) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(c, user.contextTimeout)
	defer cancel()

	res, err := user.userRepo.Update(ctx, m, id)
	if err != nil {
		return res, err
	}

	return res, nil
}
