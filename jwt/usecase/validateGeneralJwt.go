package usecase

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (h *JwtUsecase) SetJwtGeneral(g *echo.Group) {
	secret := h.Config.GetString("jwt.secret")

	// validate jwt token
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(secret),
	}))

	// validate payload related with admin type of token
	g.Use(h.ValidateGeneralJwt)
}

//ValidateGeneralJwt Use this method to Get Data Either ADMIN or MERCHANT
func (h *JwtUsecase) ValidateGeneralJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["is_admin"] == true {
				return next(c)
			} else {
				mid, ok := claims["jti"].(string)
				if !ok {
					return echo.NewHTTPError(http.StatusForbidden, "something wrong with your token id")
				}
				ctx := context.TODO()
				user, err := h.getOneUser(ctx, mid)
				if err != nil {
					return echo.NewHTTPError(http.StatusForbidden, "forbidden")
				}

				c.Set("user", user)
			}
			return next(c)
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}
