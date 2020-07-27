package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	_userHttp "github.com/bxcodec/go-clean-arch/user/delivery/http"
	_userRepo "github.com/bxcodec/go-clean-arch/user/repository/mongo"
	_userUcase "github.com/bxcodec/go-clean-arch/user/usecase"

	_catHttp "github.com/bxcodec/go-clean-arch/cat/delivery/http"
	_catRepo "github.com/bxcodec/go-clean-arch/cat/repository/mongo"
	_catUcase "github.com/bxcodec/go-clean-arch/cat/usecase"

	_jwt "github.com/bxcodec/go-clean-arch/jwt/usecase"

	_loginHttp "github.com/bxcodec/go-clean-arch/login/delivery/http"
	_loginUsecase "github.com/bxcodec/go-clean-arch/login/usecase"

	"github.com/bxcodec/go-clean-arch/bootstrap"
)

func main() {
	e := echo.New()

	timeoutContext := time.Duration(bootstrap.App.Config.GetInt("context.timeout")) * time.Second

	database := bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongo.name"))

	userRepo := _userRepo.NewMongoRepository(database)
	usrUsecase := _userUcase.NewUserUsecase(userRepo, timeoutContext)
	_userHttp.NewUserHandler(e, usrUsecase)

	jwt := _jwt.NewJwtUsecase(userRepo, timeoutContext, bootstrap.App.Config)
	userJwt := e.Group("")
	jwt.SetJwtUser(userJwt)
	adminJwt := e.Group("")
	jwt.SetJwtUser(adminJwt)
	generalJwt := e.Group("")
	jwt.SetJwtUser(generalJwt)

	//Handle For login endpoint
	loginUsecase := _loginUsecase.NewLoginUsecase(userRepo, timeoutContext)
	_loginHttp.NewLoginHandler(e, loginUsecase, bootstrap.App.Config)

	catRepo := _catRepo.NewMongoRepository(database)
	catUsecase := _catUcase.NewCatUsecase(catRepo, timeoutContext)

	_catHttp.NewCatHandler(userJwt, catUsecase)

	appPort := fmt.Sprintf(":%s", bootstrap.App.Config.GetString("server.address"))
	log.Fatal(e.Start(appPort))
}
