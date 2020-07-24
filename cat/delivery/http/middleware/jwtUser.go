package middleware

import (
	_hUsecase "github.com/bxcodec/go-clean-arch/helpers/usecase"
	"github.com/labstack/echo"
)

func RequireJwtUser(g *echo.Group) {
	h := _hUsecase.NewHelpersUsecase()
	h.SetJwtUser(g)
}
