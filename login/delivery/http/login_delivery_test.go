package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	loginHttp "github.com/bxcodec/go-clean-arch/login/delivery/http"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
)

func TestLoginUser(t *testing.T) {
	mockUCase := new(mocks.LoginUsecase)
	mockUser := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "favian",
		Password:  "password",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	mockLogin := &domain.Login{
		Username: "vian",
		Password: "password",
	}

	config := viper.New()

	tempMockLogin := mockLogin
	j, err := json.Marshal(tempMockLogin)
	assert.NoError(t, err)

	mockUCase.On("GetUser", mock.Anything, mock.Anything, mock.Anything).Return(mockUser, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/login", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")

	handler := loginHttp.LoginHandler{
		LoginUsecase: mockUCase,
		Config:       config,
	}

	err = handler.CreateJwtUser(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestLoginUserFailed(t *testing.T) {
	mockUCase := new(mocks.LoginUsecase)
	mockUser := &domain.User{
		ID:        primitive.NewObjectID(),
		Name:      "vian",
		Username:  "favian",
		Password:  "password",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	mockLogin := &domain.Login{
		Username: "vian",
		Password: "password",
	}

	config := viper.New()

	tempMockLogin := mockLogin
	j, err := json.Marshal(tempMockLogin)
	assert.NoError(t, err)

	mockUCase.On("GetUser", mock.Anything, mock.Anything, mock.Anything).Return(mockUser, errors.New("Unexpected"))

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/login", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")

	handler := loginHttp.LoginHandler{
		LoginUsecase: mockUCase,
		Config:       config,
	}

	err = handler.CreateJwtUser(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestLoginAdmin(t *testing.T) {
	mockUCase := new(mocks.LoginUsecase)
	mockLogin := &domain.Login{
		Username: "vian",
		Password: "password",
	}

	config := viper.New()
	config.Set("admin.username", "vian")
	config.Set("admin.password", "password")

	tempMockLogin := mockLogin
	j, err := json.Marshal(tempMockLogin)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/login/admin", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login/admin")

	handler := loginHttp.LoginHandler{
		LoginUsecase: mockUCase,
		Config:       config,
	}

	err = handler.CreateJwtAdmin(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestLoginAdminFailed(t *testing.T) {
	mockUCase := new(mocks.LoginUsecase)
	mockLogin := &domain.Login{
		Username: "failed",
		Password: "failed",
	}

	config := viper.New()
	config.Set("admin.username", "vian")
	config.Set("admin.password", "password")

	tempMockLogin := mockLogin
	j, err := json.Marshal(tempMockLogin)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/login/admin", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login/admin")

	handler := loginHttp.LoginHandler{
		LoginUsecase: mockUCase,
		Config:       config,
	}

	err = handler.CreateJwtAdmin(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	mockUCase.AssertExpectations(t)
}
