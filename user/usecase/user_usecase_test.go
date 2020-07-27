package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
	ucase "github.com/bxcodec/go-clean-arch/user/usecase"
)

func TestInsertOne(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &domain.User{
		Name:     "vian",
		Username: "favian",
		Password: "password",
	}
	mockEmptyUser := &domain.User{}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockUser, nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)

		a, err := u.InsertOne(context.TODO(), mockUser)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockEmptyUser, errors.New("Unexpected")).Once()

		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)

		a, err := u.InsertOne(context.TODO(), mockUser)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyUser, a)

		mockUserRepo.AssertExpectations(t)
	})

}

func TestFindOne(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "favian",
		Password:  "password",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	mockEmptyUser := &domain.User{}
	UserID := mock.Anything

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindOne", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)

		a, err := u.FindOne(context.TODO(), UserID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("FindOne", mock.Anything, mock.Anything).Return(mockEmptyUser, errors.New("Unexpected")).Once()

		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)

		a, err := u.FindOne(context.TODO(), UserID)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyUser, a)

		mockUserRepo.AssertExpectations(t)
	})

}

func TestGetAllWithPage(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockArrayUser := []domain.User{
		domain.User{
			ID:        primitive.NewObjectID(),
			Name:      "vian",
			Username:  "favian",
			Password:  "password",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
		domain.User{
			ID:        primitive.NewObjectID(),
			Name:      "vian",
			Username:  "testt",
			Password:  "test",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}
	mockEmptyUser := []domain.User{domain.User{}}
	p := int64(1)
	rp := int64(25)
	filter := mock.Anything
	setsort := mock.Anything
	count := int64(25)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetAllWithPage", mock.Anything, mock.AnythingOfType("int64"), mock.Anything).Return(mockArrayUser, count, nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)

		a, count, err := u.GetAllWithPage(context.TODO(), rp, p, filter, setsort)

		assert.NoError(t, err)
		assert.NotNil(t, a)
		assert.Equal(t, count, int64(25))

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("GetAllWithPage", mock.Anything, mock.AnythingOfType("int64"), mock.Anything).Return(mockEmptyUser, count, errors.New("Unexpected")).Once()

		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)

		a, count, err := u.GetAllWithPage(context.TODO(), rp, p, filter, setsort)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyUser, a)
		assert.Equal(t, count, int64(25))

		mockUserRepo.AssertExpectations(t)
	})

}

func TestUpdateOne(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "testt",
		Password:  "test",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	mockEmptyUser := &domain.User{}
	UserID := mock.Anything

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("UpdateOne", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)

		a, err := u.UpdateOne(context.TODO(), mockUser, UserID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("UpdateOne", mock.Anything, mock.Anything).Return(mockEmptyUser, errors.New("Unexpected")).Once()

		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)

		a, err := u.UpdateOne(context.TODO(), mockUser, UserID)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyUser, a)

		mockUserRepo.AssertExpectations(t)
	})

}
