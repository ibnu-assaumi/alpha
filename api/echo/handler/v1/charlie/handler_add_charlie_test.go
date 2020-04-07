package charlie

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Bhinneka/alpha/api/lib/response"
	charlieUseCase "github.com/Bhinneka/alpha/api/service/usecase/v1/charlie/mocks"
)

func Test_handler_addCharlie(t *testing.T) {
	t.Run("POSITIVE_ADD_CHARLIE", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/v1/charlie", strings.NewReader(`{ "charlieName" : "test"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		res := response.Success(response.Meta{}, []string{}, "success")

		mockUseCase := new(charlieUseCase.UseCase)
		mockUseCase.On("AddCharlie", mock.Anything, mock.Anything).Return(res)
		h := &handler{
			usecaseCharlie: mockUseCase,
		}

		if assert.NoError(t, h.addCharlie(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("NEGATIVE_ADD_CHARLIE", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/v1/charlie", strings.NewReader(`{ "charlieName" : 123}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		res := response.BadRequest("error")

		mockUseCase := new(charlieUseCase.UseCase)
		mockUseCase.On("AddCharlie", mock.Anything, mock.Anything).Return(res)
		h := &handler{
			usecaseCharlie: mockUseCase,
		}

		if assert.NoError(t, h.addCharlie(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
