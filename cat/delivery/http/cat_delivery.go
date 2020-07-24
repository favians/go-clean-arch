package http

import (
	"context"
	"math"
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/cat/delivery/http/middleware"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CatHandler struct {
	CatUsecase domain.CatUsecase
}

func NewCatHandler(e *echo.Echo, uu domain.CatUsecase) {
	handler := &CatHandler{
		CatUsecase: uu,
	}
	g := e.Group("")
	middleware.RequireJwtUser(g)
	g.POST("/cat", handler.Store)
	g.GET("/cat", handler.GetOne)
	g.GET("/cats", handler.GetAll)
	g.PUT("/cat", handler.Update)
}

func isRequestValid(m *domain.Cat) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cat *CatHandler) Store(c echo.Context) error {
	var (
		ct  domain.Cat
		err error
	)
	user := c.Get("user")
	token := user.(*domain.User)

	err = c.Bind(&ct)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&ct); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ct.UserID = token.ID
	result, err := cat.CatUsecase.Store(ctx, &ct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (cat *CatHandler) GetOne(c echo.Context) error {

	id := c.QueryParam("id")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := cat.CatUsecase.GetOne(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (cat *CatHandler) GetAll(c echo.Context) error {

	type Response struct {
		Total       int64        `json:"total"`
		PerPage     int64        `json:"per_page"`
		CurrentPage int64        `json:"current_page"`
		LastPage    int64        `json:"last_page"`
		From        int64        `json:"from"`
		To          int64        `json:"to"`
		Cat         []domain.Cat `json:"cats"`
	}

	var (
		res   []domain.Cat
		count int64
	)

	rp, err := strconv.ParseInt(c.QueryParam("rp"), 10, 64)
	if err != nil {
		rp = 25
	}

	page, err := strconv.ParseInt(c.QueryParam("p"), 10, 64)
	if err != nil {
		page = 1
	}

	filters := bson.D{{"name", primitive.Regex{Pattern: ".*" + c.QueryParam("name") + ".*", Options: "i"}}}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, count, err = cat.CatUsecase.GetAllWithPage(ctx, rp, page, filters, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result := Response{
		Total:       count,
		PerPage:     rp,
		CurrentPage: page,
		LastPage:    int64(math.Ceil(float64(count) / float64(rp))),
		From:        page*rp - rp + 1,
		To:          page * rp,
		Cat:         res,
	}

	return c.JSON(http.StatusOK, result)
}

func (cat *CatHandler) Update(c echo.Context) error {

	id := c.QueryParam("id")

	var (
		ct  domain.Cat
		err error
	)

	err = c.Bind(&ct)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := cat.CatUsecase.Update(ctx, &ct, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
