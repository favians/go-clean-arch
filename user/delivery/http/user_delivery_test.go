package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	userHttp "github.com/bxcodec/go-clean-arch/user/delivery/http"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
)

func TestStore(t *testing.T) {
	mockUCase := new(mocks.UserUsecase)
	mockUserRequest := &domain.User{
		Name:     "vian",
		Username: "favian",
		Password: "password",
	}
	mockUserResponse := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "favian",
		Password:  "password",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	tempMockUser := mockUserRequest
	j, err := json.Marshal(tempMockUser)
	assert.NoError(t, err)

	mockUCase.On("Store", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockUserResponse, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/user", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	handler := userHttp.UserHandler{
		UsrUsecase: mockUCase,
	}

	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestStoreFailed(t *testing.T) {
	mockUCase := new(mocks.UserUsecase)
	mockUserRequest := &domain.User{
		Name:     "vian",
		Username: "favian",
		Password: "password",
	}
	mockUserResponse := &domain.User{}
	tempMockUser := mockUserRequest
	j, err := json.Marshal(tempMockUser)
	assert.NoError(t, err)

	mockUCase.On("Store", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockUserResponse, errors.New("Unexpected"))

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/user", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	handler := userHttp.UserHandler{
		UsrUsecase: mockUCase,
	}

	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetOne(t *testing.T) {
	mockUCase := new(mocks.UserUsecase)
	mockUser := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "favian",
		Password:  "password",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	UserID := mock.Anything

	mockUCase.On("GetOne", mock.Anything, UserID).Return(mockUser, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/user", nil)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	handler := userHttp.UserHandler{
		UsrUsecase: mockUCase,
	}

	err = handler.GetOne(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetOneFailed(t *testing.T) {
	mockUCase := new(mocks.UserUsecase)
	mockUserFailed := &domain.User{}
	UserID := mock.Anything

	mockUCase.On("GetOne", mock.Anything, UserID).Return(mockUserFailed, errors.New("Unexpected"))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/user", nil)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	handler := userHttp.UserHandler{
		UsrUsecase: mockUCase,
	}

	err = handler.GetOne(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockUCase := new(mocks.UserUsecase)
	mockUserRequest := &domain.User{
		Name:     "vian",
		Username: "favian",
		Password: "password",
	}
	mockUserResponse := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "favian",
		Password:  "password",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	UserID := mock.Anything

	tempMockUser := mockUserRequest
	j, err := json.Marshal(tempMockUser)
	assert.NoError(t, err)

	mockUCase.On("Update", mock.Anything, mock.AnythingOfType("*domain.User"), UserID).Return(mockUserResponse, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/user", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	handler := userHttp.UserHandler{
		UsrUsecase: mockUCase,
	}

	err = handler.Update(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUpdateFailed(t *testing.T) {
	mockUCase := new(mocks.UserUsecase)
	mockUserRequest := &domain.User{
		Name:     "vian",
		Username: "favian",
		Password: "password",
	}
	mockUserResponse := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "favian",
		Password:  "password",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	UserID := mock.Anything

	tempMockUser := mockUserRequest
	j, err := json.Marshal(tempMockUser)
	assert.NoError(t, err)

	mockUCase.On("Update", mock.Anything, mock.AnythingOfType("*domain.User"), UserID).Return(mockUserResponse, errors.New("Unexpected"))

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/user", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	handler := userHttp.UserHandler{
		UsrUsecase: mockUCase,
	}

	err = handler.Update(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAll(t *testing.T) {
	mockUCase := new(mocks.UserUsecase)
	mockUser := []domain.User{
		domain.User{
			ID:        primitive.NewObjectID(),
			Name:      "vian",
			Username:  "favian",
			Password:  "password",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}
	rp := mock.Anything
	p := mock.Anything
	filters := mock.Anything
	setsort := mock.Anything
	count := int64(25)

	mockUCase.On("GetAllWithPage", mock.Anything, rp, p, filters, setsort).Return(mockUser, count, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/users", nil)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	handler := userHttp.UserHandler{
		UsrUsecase: mockUCase,
	}

	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAllFailed(t *testing.T) {
	mockUCase := new(mocks.UserUsecase)
	mockUserFailed := []domain.User{domain.User{}, domain.User{}}
	rp := mock.Anything
	p := mock.Anything
	filters := mock.Anything
	setsort := mock.Anything
	count := int64(25)

	mockUCase.On("GetAllWithPage", mock.Anything, rp, p, filters, setsort).Return(mockUserFailed, count, errors.New("Unexpected"))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/users", nil)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	handler := userHttp.UserHandler{
		UsrUsecase: mockUCase,
	}

	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}
