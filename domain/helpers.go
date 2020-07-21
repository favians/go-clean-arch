package domain

import "github.com/labstack/echo"

type Helpers struct{}

type HelpersUsecase interface {
	SetJwtAdmin(g *echo.Group)
	SetJwtUser(g *echo.Group)
	SetJwtGeneral(g *echo.Group)
}
