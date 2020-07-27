package usecase_test

import (
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/bxcodec/go-clean-arch/domain/mocks"
	ucase "github.com/bxcodec/go-clean-arch/jwt/usecase"
)

func TestSetJwtAdmin(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	// mockUser := &domain.User{
	// 	CreatedAt: time.Now(),
	// 	ID:        primitive.NewObjectID(),
	// 	Password:  "password",
	// 	Username:  "username",
	// 	UpdatedAt: time.Now(),
	// }
	// userID := mock.Anything
	config := new(viper.Viper)

	t.Run("success", func(t *testing.T) {
		// mockUserRepo.On("GetOne", mock.Anything, userID).Return(mockUser, nil).Once()

		u := ucase.NewJwtUsecase(mockUserRepo, time.Second*2, config)

		e := echo.New()
		userJwt := e.Group("")

		u.SetJwtAdmin(userJwt)

		mockUserRepo.AssertExpectations(t)
	})
}
