package usecase

import (
	"context"
	"net/http"

	"github.com/bxcodec/go-clean-arch/bootstrap"
	"github.com/bxcodec/go-clean-arch/domain"
	_userRepo "github.com/bxcodec/go-clean-arch/user/repository/mongo"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//SetJwtUser Set Only JWT for For User
func (h *jwtUsecase) SetJwtUser(g *echo.Group) {

	secret := bootstrap.App.Config.GetString("jwt.secret")
	// validate jwt token
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(secret),
	}))

	// validate payload related with user
	g.Use(h.validateJwtUser)
}

func (h *jwtUsecase) validateJwtUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
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

			return next(c)
		}

		return echo.NewHTTPError(http.StatusForbidden, "invalid token")
	}
}

func findUser(id string) (*domain.User, error) {
	database := bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongo.name"))
	userRepo := _userRepo.NewMongoRepository(database)

	ctx := context.TODO()

	res, err := userRepo.GetOne(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}
