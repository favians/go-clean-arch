package http_test

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	catHttp "github.com/bxcodec/go-clean-arch/cat/delivery/http"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
)

func TestInsertOne(t *testing.T) {
	mockUCase := new(mocks.CatUsecase)
	mockCatRequest := &domain.Cat{
		Name:    "blacky",
		Legs:    4,
		Species: "kucing item",
	}
	mockCatResponse := &domain.Cat{
		ID:        primitive.NewObjectID(),
		Name:      "blacky",
		Legs:      4,
		Species:   "kucing item",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		UserID:    primitive.NewObjectID(),
	}
	mockUser := &domain.User{
		ID: primitive.NewObjectID(),
	}
	tempMockCat := mockCatRequest
	j, err := json.Marshal(tempMockCat)
	assert.NoError(t, err)

	mockUCase.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.Cat")).Return(mockCatResponse, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/cat", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cat")
	c.Set("user", mockUser)

	handler := catHttp.CatHandler{
		CatUsecase: mockUCase,
	}

	err = handler.InsertOne(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestInsertOneFailed(t *testing.T) {
	mockUCase := new(mocks.CatUsecase)
	mockCatRequest := &domain.Cat{
		Name:    "blacky",
		Legs:    4,
		Species: "kucing item",
	}
	mockCatResponse := &domain.Cat{}
	mockUser := &domain.User{
		ID: primitive.NewObjectID(),
	}
	tempMockCat := mockCatRequest
	j, err := json.Marshal(tempMockCat)
	assert.NoError(t, err)

	mockUCase.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.Cat")).Return(mockCatResponse, errors.New("Unexpected"))

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/cat", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cat")
	c.Set("user", mockUser)

	handler := catHttp.CatHandler{
		CatUsecase: mockUCase,
	}

	err = handler.InsertOne(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFindOne(t *testing.T) {
	mockUCase := new(mocks.CatUsecase)
	mockCat := &domain.Cat{
		ID:        primitive.NewObjectID(),
		Name:      "blacky",
		Legs:      4,
		Species:   "kucing item",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		UserID:    primitive.NewObjectID(),
	}
	CatID := mock.Anything

	mockUCase.On("FindOne", mock.Anything, CatID).Return(mockCat, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/cat", nil)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cat")

	handler := catHttp.CatHandler{
		CatUsecase: mockUCase,
	}
	log.Println(handler)

	err = handler.FindOne(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFindOneFailed(t *testing.T) {
	mockUCase := new(mocks.CatUsecase)
	mockCatFailed := &domain.Cat{}
	CatID := mock.Anything

	mockUCase.On("FindOne", mock.Anything, CatID).Return(mockCatFailed, errors.New("Unexpected"))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/cat", nil)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cat")

	handler := catHttp.CatHandler{
		CatUsecase: mockUCase,
	}
	log.Println(handler)

	err = handler.FindOne(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUpdateOne(t *testing.T) {
	mockUCase := new(mocks.CatUsecase)
	mockCatRequest := &domain.Cat{
		Name:    "blacky",
		Legs:    4,
		Species: "kucing item",
	}
	mockCatResponse := &domain.Cat{
		ID:        primitive.NewObjectID(),
		Name:      "blacky",
		Legs:      4,
		Species:   "kucing item",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		UserID:    primitive.NewObjectID(),
	}
	mockUser := &domain.User{
		ID: primitive.NewObjectID(),
	}
	CatID := mock.Anything

	tempMockCat := mockCatRequest
	j, err := json.Marshal(tempMockCat)
	assert.NoError(t, err)

	mockUCase.On("UpdateOne", mock.Anything, mock.AnythingOfType("*domain.Cat"), CatID).Return(mockCatResponse, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/cat", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cat")
	c.Set("user", mockUser)

	handler := catHttp.CatHandler{
		CatUsecase: mockUCase,
	}

	err = handler.UpdateOne(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUpdateOneFailed(t *testing.T) {
	mockUCase := new(mocks.CatUsecase)
	mockCatRequest := &domain.Cat{
		Name:    "blacky",
		Legs:    4,
		Species: "kucing item",
	}
	mockCatResponse := &domain.Cat{}
	mockUser := &domain.User{
		ID: primitive.NewObjectID(),
	}
	CatID := mock.Anything

	tempMockCat := mockCatRequest
	j, err := json.Marshal(tempMockCat)
	assert.NoError(t, err)

	mockUCase.On("UpdateOne", mock.Anything, mock.AnythingOfType("*domain.Cat"), CatID).Return(mockCatResponse, errors.New("Unexpected"))

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/cat", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cat")
	c.Set("user", mockUser)

	handler := catHttp.CatHandler{
		CatUsecase: mockUCase,
	}

	err = handler.UpdateOne(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAll(t *testing.T) {
	mockUCase := new(mocks.CatUsecase)
	mockCat := []domain.Cat{
		domain.Cat{
			ID:        primitive.NewObjectID(),
			Name:      "blacky",
			Legs:      4,
			Species:   "kucing item",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
			UserID:    primitive.NewObjectID(),
		},
	}
	rp := mock.Anything
	p := mock.Anything
	filters := mock.Anything
	setsort := mock.Anything
	count := int64(25)

	mockUCase.On("GetAllWithPage", mock.Anything, rp, p, filters, setsort).Return(mockCat, count, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/cats", nil)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cats")

	handler := catHttp.CatHandler{
		CatUsecase: mockUCase,
	}
	log.Println(handler)

	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAllFailed(t *testing.T) {
	mockUCase := new(mocks.CatUsecase)
	mockCatFailed := []domain.Cat{domain.Cat{}, domain.Cat{}}
	rp := mock.Anything
	p := mock.Anything
	filters := mock.Anything
	setsort := mock.Anything
	count := int64(25)

	mockUCase.On("GetAllWithPage", mock.Anything, rp, p, filters, setsort).Return(mockCatFailed, count, errors.New("Unexpected"))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/cats", nil)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cats")

	handler := catHttp.CatHandler{
		CatUsecase: mockUCase,
	}
	log.Println(handler)

	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}
