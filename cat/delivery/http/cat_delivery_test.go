package http_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	articleHttp "github.com/bxcodec/go-clean-arch/cat/delivery/http"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
)

func TestGetOne(t *testing.T) {
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

	mockUCase.On("GetOne", mock.Anything, CatID).Return(mockCat, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/cat", nil)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cat")

	handler := articleHttp.CatHandler{
		CatUsecase: mockUCase,
	}
	log.Println(handler)

	err = handler.GetOne(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}
