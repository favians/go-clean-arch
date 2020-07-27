package usecase_test

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
	ucase "github.com/bxcodec/go-clean-arch/jwt/usecase"
)

func TestGetOneUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &domain.User{
		CreatedAt: time.Now(),
		ID:        primitive.NewObjectID(),
		Password:  "password",
		Username:  "username",
		UpdatedAt: time.Now(),
	}
	userID := mock.Anything
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetOne", mock.Anything, userID).Return(mockUser, nil).Once()

		u := ucase.NewJwtUsecase(mockUserRepo, time.Second*2)

		e := echo.New()
		userJwt := e.Group("")

		u.SetJwtAdmin(userJwt)

		mockUserRepo.AssertExpectations(t)
	})
}
