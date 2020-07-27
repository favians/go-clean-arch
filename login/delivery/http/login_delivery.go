package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type LoginHandler struct {
	LoginUsecase domain.LoginUsecase
	Config       *viper.Viper
}

func NewLoginHandler(e *echo.Echo, lu domain.LoginUsecase, config *viper.Viper) {
	handler := &LoginHandler{
		LoginUsecase: lu,
		Config:       config,
	}
	e.POST("/login/admin", handler.CreateJwtAdmin)
	e.POST("/login", handler.CreateJwtUser)
}

func isRequestValid(m *domain.Login) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (login *LoginHandler) CreateJwtUser(c echo.Context) error {

	var (
		err          error
		token        string
		loginPayload domain.Login
	)

	err = c.Bind(&loginPayload)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&loginPayload); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := login.LoginUsecase.GetUser(ctx, loginPayload.Username, loginPayload.Password)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Your username or password were wrong")
	}

	lifetime, err := strconv.ParseInt(login.Config.GetString("jwt.lifetime"), 10, 64)
	if err != nil {
		lifetime = 60
	}

	secret := login.Config.GetString("jwt.secret")
	token, err = createJwtToken(res.ID.Hex(), "user", lifetime, secret)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token":      token,
		"expires_in": lifetime,
	})

}

func (login *LoginHandler) CreateJwtAdmin(c echo.Context) error {

	var (
		err          error
		token        string
		loginPayload domain.Login
	)

	err = c.Bind(&loginPayload)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&loginPayload); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	adminUsername, adminPassword := login.Config.GetString("admin.username"), login.Config.GetString("admin.password")

	if loginPayload.Username == adminUsername && loginPayload.Password == adminPassword {
		// create jwt token
		lifetime, err := strconv.ParseInt(login.Config.GetString("jwt.lifetime"), 10, 64)
		if err != nil {
			lifetime = 60
		}

		secret := login.Config.GetString("jwt.secret")
		token, err = createJwtToken(adminUsername, "admin", lifetime, secret)
		if err != nil {
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"expires_in": lifetime,
			"token":      token,
		})
	}

	return c.String(http.StatusUnauthorized, "Your username or password were wrong")
}

func createJwtToken(uname string, jtype string, lifetime int64, secret string) (string, error) {

	type JwtClaims struct {
		Name    string `json:"name"`
		IsAdmin bool   `json:"is_admin"`
		jwt.StandardClaims
	}

	getLifeTime := lifetime
	getTime := time.Duration(getLifeTime)

	var (
		claim    JwtClaims
		lifeTime int64 = time.Now().Add(getTime * time.Minute).Unix()
	)

	if jtype == "admin" {
		claim = JwtClaims{
			uname,
			true,
			jwt.StandardClaims{
				Id:        uname,
				ExpiresAt: lifeTime,
			},
		}
	} else {
		claim = JwtClaims{
			uname,
			false,
			jwt.StandardClaims{
				Id:        uname,
				ExpiresAt: lifeTime,
			},
		}
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	token, err := rawToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
