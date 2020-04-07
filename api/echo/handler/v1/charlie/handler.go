package charlie

import (
	"github.com/labstack/echo/v4"

	usecaseCharlie "github.com/Bhinneka/alpha/api/service/usecase/v1/charlie"
)

type handler struct {
	usecaseCharlie usecaseCharlie.UseCase
}

// Handler : charlie rest http handler
type Handler interface {
	Mount(group *echo.Group)
}

// NewHandler : new charlie rest http handler
func NewHandler() Handler {
	return handler{
		usecaseCharlie: usecaseCharlie.NewUseCase(),
	}
}

func (impl handler) Mount(group *echo.Group) {
	group.GET("", impl.getCharlie)
	group.POST("", impl.addCharlie)
	group.PUT("", impl.updateCharlie)
	group.DELETE("", impl.deleteCharlie)
}
