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
	ucase "github.com/bxcodec/go-clean-arch/login/usecase"
)

func TestGetUser(t *testing.T) {
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
	username := "vian"
	password := "pasword"

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByCredential", mock.Anything, mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		u := ucase.NewLoginUsecase(mockUserRepo, time.Second*2)

		a, err := u.GetUser(context.TODO(), username, password)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("GetByCredential", mock.Anything, mock.Anything, mock.Anything).Return(mockEmptyUser, errors.New("Unexpected")).Once()

		u := ucase.NewLoginUsecase(mockUserRepo, time.Second*2)

		a, err := u.GetUser(context.TODO(), username, password)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyUser, a)

		mockUserRepo.AssertExpectations(t)
	})

}
