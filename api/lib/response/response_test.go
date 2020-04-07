package response

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadRequest(t *testing.T) {
	response := BadRequest("invalid request")
	assert.NotNil(t, response)
	assert.Equal(t, false, response.Success)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestInternalServerError(t *testing.T) {
	response := InternalServerError()
	assert.NotNil(t, response)
	assert.Equal(t, false, response.Success)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func TestNotFound(t *testing.T) {
	response := NotFound("page / record not found")
	assert.NotNil(t, response)
	assert.Equal(t, false, response.Success)
	assert.Equal(t, http.StatusNotFound, response.Code)
}

func TestSuccess(t *testing.T) {
	meta := Meta{
		Limit:     10,
		TotalData: 10,
		TotalPage: 1,
	}
	response := Success(meta, []string{}, "success")
	assert.NotNil(t, response)
}

func TestUnauthorized(t *testing.T) {
	response := Unauthorized()
	assert.NotNil(t, response)
	assert.Equal(t, false, response.Success)
	assert.Equal(t, http.StatusUnauthorized, response.Code)
}
