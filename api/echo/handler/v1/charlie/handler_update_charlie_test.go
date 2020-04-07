package charlie

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Bhinneka/alpha/api/lib/response"
	charlieUseCase "github.com/Bhinneka/alpha/api/service/usecase/v1/charlie/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_handler_updateCharlie(t *testing.T) {
	t.Run("POSITIVE_UPDATE_CHARLIE", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/v1/charlie", strings.NewReader(`{ "carlieID" : 1, "charlieName" : "test"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		res := response.Success(response.Meta{}, []string{}, "success")

		mockUseCase := new(charlieUseCase.UseCase)
		mockUseCase.On("UpdateCharlie", mock.Anything, mock.Anything).Return(res)
		h := &handler{
			usecaseCharlie: mockUseCase,
		}

		if assert.NoError(t, h.updateCharlie(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("NEGATIVE_UPDATE_CHARLIE", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/v1/charlie", strings.NewReader(`{ "charlieID" : 1, "charlieName" : 123}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		res := response.BadRequest("error")

		mockUseCase := new(charlieUseCase.UseCase)
		mockUseCase.On("UpdateCharlie", mock.Anything, mock.Anything).Return(res)
		h := &handler{
			usecaseCharlie: mockUseCase,
		}

		if assert.NoError(t, h.updateCharlie(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
