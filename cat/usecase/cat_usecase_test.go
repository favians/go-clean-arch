package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	ucase "github.com/bxcodec/go-clean-arch/cat/usecase"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
)

func TestInsertOne(t *testing.T) {
	mockCatRepo := new(mocks.CatRepository)
	mockCat := &domain.Cat{
		Name:    "blacky",
		Legs:    4,
		Species: "kucing item",
		UserID:  primitive.NewObjectID(),
	}
	mockEmptyCat := &domain.Cat{}

	t.Run("success", func(t *testing.T) {
		mockCatRepo.On("InsertOne", mock.Anything, mock.Anything).Return(mockCat, nil).Once()
		u := ucase.NewCatUsecase(mockCatRepo, time.Second*2)

		a, err := u.InsertOne(context.TODO(), mockCat)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockCatRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCatRepo.On("InsertOne", mock.Anything, mock.Anything).Return(mockEmptyCat, errors.New("Unexpected")).Once()

		u := ucase.NewCatUsecase(mockCatRepo, time.Second*2)

		a, err := u.InsertOne(context.TODO(), mockCat)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyCat, a)

		mockCatRepo.AssertExpectations(t)
	})

}

func TestFindOne(t *testing.T) {
	mockCatRepo := new(mocks.CatRepository)
	mockCat := &domain.Cat{
		ID:        primitive.NewObjectID(),
		Name:      "blacky",
		Legs:      4,
		Species:   "kucing item",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		UserID:    primitive.NewObjectID(),
	}
	mockEmptyCat := &domain.Cat{}
	CatID := mock.Anything

	t.Run("success", func(t *testing.T) {
		mockCatRepo.On("FindOne", mock.Anything, mock.Anything).Return(mockCat, nil).Once()
		u := ucase.NewCatUsecase(mockCatRepo, time.Second*2)

		a, err := u.FindOne(context.TODO(), CatID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockCatRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCatRepo.On("FindOne", mock.Anything, mock.Anything).Return(mockEmptyCat, errors.New("Unexpected")).Once()

		u := ucase.NewCatUsecase(mockCatRepo, time.Second*2)

		a, err := u.FindOne(context.TODO(), CatID)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyCat, a)

		mockCatRepo.AssertExpectations(t)
	})

}

func TestGetAllWithPage(t *testing.T) {
	mockCatRepo := new(mocks.CatRepository)
	mockArrayCat := []domain.Cat{
		domain.Cat{
			Name:    "blacky",
			Legs:    4,
			Species: "kucing item",
			UserID:  primitive.NewObjectID(),
		},
		domain.Cat{
			Name:    "whity",
			Legs:    4,
			Species: "anggora",
			UserID:  primitive.NewObjectID(),
		},
	}
	mockEmptyCat := []domain.Cat{domain.Cat{}}
	p := int64(1)
	rp := int64(25)
	filter := mock.Anything
	setsort := mock.Anything
	count := int64(25)

	t.Run("success", func(t *testing.T) {
		mockCatRepo.On("GetAllWithPage", mock.Anything, mock.AnythingOfType("int64"), mock.Anything).Return(mockArrayCat, count, nil).Once()
		u := ucase.NewCatUsecase(mockCatRepo, time.Second*2)

		a, count, err := u.GetAllWithPage(context.TODO(), rp, p, filter, setsort)

		assert.NoError(t, err)
		assert.NotNil(t, a)
		assert.Equal(t, count, int64(25))

		mockCatRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCatRepo.On("GetAllWithPage", mock.Anything, mock.AnythingOfType("int64"), mock.Anything).Return(mockEmptyCat, count, errors.New("Unexpected")).Once()

		u := ucase.NewCatUsecase(mockCatRepo, time.Second*2)

		a, count, err := u.GetAllWithPage(context.TODO(), rp, p, filter, setsort)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyCat, a)
		assert.Equal(t, count, int64(25))

		mockCatRepo.AssertExpectations(t)
	})

}

func TestUpdateOne(t *testing.T) {
	mockCatRepo := new(mocks.CatRepository)
	mockCat := &domain.Cat{
		ID:        primitive.NewObjectID(),
		Name:      "blacky",
		Legs:      4,
		Species:   "kucing item",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		UserID:    primitive.NewObjectID(),
	}
	mockEmptyCat := &domain.Cat{}
	CatID := mock.Anything

	t.Run("success", func(t *testing.T) {
		mockCatRepo.On("UpdateOne", mock.Anything, mock.Anything).Return(mockCat, nil).Once()
		u := ucase.NewCatUsecase(mockCatRepo, time.Second*2)

		a, err := u.UpdateOne(context.TODO(), mockCat, CatID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockCatRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCatRepo.On("UpdateOne", mock.Anything, mock.Anything).Return(mockEmptyCat, errors.New("Unexpected")).Once()

		u := ucase.NewCatUsecase(mockCatRepo, time.Second*2)

		a, err := u.UpdateOne(context.TODO(), mockCat, CatID)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyCat, a)

		mockCatRepo.AssertExpectations(t)
	})

}
