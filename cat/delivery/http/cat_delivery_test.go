package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/bxcodec/faker"
	catHttp "github.com/bxcodec/go-clean-arch/cat/delivery/http"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
)

func TestGetOne(t *testing.T) {
	var mockCat domain.Cat
	err := faker.FakeData(&mockCat)
	assert.NoError(t, err)

	mockUCase := new(mocks.CatUsecase)
	CatID := mock.Anything

	t.Run("success", func(t *testing.T) {
		mockUCase.On("GetOne", mock.Anything, CatID).Return(mockCat, nil)

		e := echo.New()
		req, err := http.NewRequest(echo.GET, "/cat", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("cat")
		handler := catHttp.CatHandler{
			CatUsecase: mockUCase,
		}
		err = handler.GetOne(c)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)
		mockUCase.AssertExpectations(t)
	})
}
