package domain

import "github.com/labstack/echo"

type Jwt struct{}

type JwtUsecase interface {
	SetJwtAdmin(g *echo.Group)
	SetJwtUser(g *echo.Group)
	SetJwtGeneral(g *echo.Group)
}
