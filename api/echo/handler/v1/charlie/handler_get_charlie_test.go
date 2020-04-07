package charlie

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Bhinneka/alpha/api/lib/response"
	charlieUseCase "github.com/Bhinneka/alpha/api/service/usecase/v1/charlie/mocks"
)

func Test_handler_getCharlie(t *testing.T) {
	t.Run("POSITIVE_GET_CHARLIE", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/charlie?charlieID=1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		res := response.Success(response.Meta{}, []string{}, "success")

		mockUseCase := new(charlieUseCase.UseCase)
		mockUseCase.On("GetCharlie", mock.Anything, mock.Anything).Return(res)
		h := &handler{
			usecaseCharlie: mockUseCase,
		}

		if assert.NoError(t, h.getCharlie(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("NEGATIVE_GET_CHARLIE", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/charlie?charlieID=test", nil)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		res := response.BadRequest("error")

		mockUseCase := new(charlieUseCase.UseCase)
		mockUseCase.On("GetCharlie", mock.Anything, mock.Anything).Return(res)
		h := &handler{
			usecaseCharlie: mockUseCase,
		}

		if assert.NoError(t, h.getCharlie(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
