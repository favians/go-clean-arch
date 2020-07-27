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
	config := new(viper.Viper)

	t.Run("success", func(t *testing.T) {
		u := ucase.NewJwtUsecase(mockUserRepo, time.Second*2, config)

		e := echo.New()
		userJwt := e.Group("")

		u.SetJwtAdmin(userJwt)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestSetJwtUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	config := new(viper.Viper)

	t.Run("success", func(t *testing.T) {
		u := ucase.NewJwtUsecase(mockUserRepo, time.Second*2, config)

		e := echo.New()
		userJwt := e.Group("")

		u.SetJwtUser(userJwt)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestSetJwtGeneral(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	config := new(viper.Viper)

	t.Run("success", func(t *testing.T) {
		u := ucase.NewJwtUsecase(mockUserRepo, time.Second*2, config)

		e := echo.New()
		userJwt := e.Group("")

		u.SetJwtGeneral(userJwt)

		mockUserRepo.AssertExpectations(t)
	})
}
