package health

import "github.com/labstack/echo/v4"

type handler struct{}

// Handler : health route handler
type Handler interface {
	Check(c echo.Context) error
}

// NewHandler : new healt route handler
func NewHandler() Handler {
	return handler{}
}
