package http

import (
	"context"
	"math"
	"net/http"
	"strconv"

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

func NewCatHandler(userJwt *echo.Group, uu domain.CatUsecase) {
	handler := &CatHandler{
		CatUsecase: uu,
	}
	userJwt.POST("/cat", handler.InsertOne)
	userJwt.GET("/cat", handler.FindOne)
	userJwt.GET("/cats", handler.GetAll)
	userJwt.PUT("/cat", handler.UpdateOne)
}

func isRequestValid(m *domain.Cat) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cat *CatHandler) InsertOne(c echo.Context) error {
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
	result, err := cat.CatUsecase.InsertOne(ctx, &ct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (cat *CatHandler) FindOne(c echo.Context) error {

	id := c.QueryParam("id")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := cat.CatUsecase.FindOne(ctx, id)
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

func (cat *CatHandler) UpdateOne(c echo.Context) error {

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

	result, err := cat.CatUsecase.UpdateOne(ctx, &ct, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
