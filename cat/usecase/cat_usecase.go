package usecase

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type catUsecase struct {
	catRepo        domain.CatRepository
	contextTimeout time.Duration
}

func NewCatUsecase(u domain.CatRepository, to time.Duration) domain.CatUsecase {
	return &catUsecase{
		catRepo:        u,
		contextTimeout: to,
	}
}

func (cat *catUsecase) Store(c context.Context, m *domain.Cat) (*domain.Cat, error) {

	ctx, cancel := context.WithTimeout(c, cat.contextTimeout)
	defer cancel()

	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	res, err := cat.catRepo.Store(ctx, m)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (cat *catUsecase) GetOne(c context.Context, id string) (*domain.Cat, error) {

	ctx, cancel := context.WithTimeout(c, cat.contextTimeout)
	defer cancel()

	res, err := cat.catRepo.GetOne(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (cat *catUsecase) GetAllWithPage(c context.Context, rp int64, p int64, filter interface{}, setsort interface{}) ([]domain.Cat, int64, error) {

	ctx, cancel := context.WithTimeout(c, cat.contextTimeout)
	defer cancel()

	res, count, err := cat.catRepo.GetAllWithPage(ctx, rp, p, filter, setsort)
	if err != nil {
		return res, count, err
	}

	return res, count, nil
}

func (cat *catUsecase) Update(c context.Context, m *domain.Cat, id string) (*domain.Cat, error) {

	ctx, cancel := context.WithTimeout(c, cat.contextTimeout)
	defer cancel()

	res, err := cat.catRepo.Update(ctx, m, id)
	if err != nil {
		return res, err
	}

	return res, nil
}
