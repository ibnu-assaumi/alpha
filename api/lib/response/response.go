package response

import (
	"net/http"
)

// Response : api response wrapper
type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Meta    Meta        `json:"meta"`
	Data    interface{} `json:"data"`
}

// Meta : response meta
type Meta struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	TotalPage int64 `json:"totalPage"`
	TotalData int64 `json:"totalData"`
}

// BadRequest : response bad request helper
func BadRequest(message interface{}) Response {
	return Response{
		Code:    http.StatusBadRequest,
		Success: false,
		Message: message,
		Data:    []string{},
	}
}

// InternalServerError : response internal server error
func InternalServerError() Response {
	return Response{
		Code:    http.StatusInternalServerError,
		Success: false,
		Message: "internal server error",
		Data:    []string{},
	}
}

// NotFound : response not found helper
func NotFound(message interface{}) Response {
	return Response{
		Code:    http.StatusNotFound,
		Success: false,
		Message: message,
		Data:    []string{},
	}
}

// Success : response success helper
func Success(meta Meta, data interface{}, message interface{}) Response {
	return Response{
		Code:    http.StatusOK,
		Success: true,
		Meta:    meta,
		Data:    data,
		Message: message,
	}
}

// Unauthorized : response unauthorized helper
func Unauthorized() Response {
	return Response{
		Code:    http.StatusUnauthorized,
		Success: false,
		Message: "unauthorized",
		Data:    []string{},
	}
}
