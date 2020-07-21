package usecase

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/bootstrap"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//SetJwtAdmin Set Only JWT for For Admin
func (h *helpersUsecase) SetJwtAdmin(g *echo.Group) {

	secret := bootstrap.App.Config.GetString("jwt.secret")

	// validate jwt token
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(secret),
	}))

	// validate payload related with admin type of token
	g.Use(h.validateJwtAdmin)
}

// validateJwtAdmin
// Middleware for validating access to Admin only resources
func (h *helpersUsecase) validateJwtAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user")
		token := user.(*jwt.Token)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["is_admin"] == true {
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}
