package charlie

import (
	"github.com/labstack/echo/v4"

	deliveryCharlie "github.com/Bhinneka/alpha/api/service/delivery/v1/charlie"
)

type handler struct {
	deliveryCharlie deliveryCharlie.Delivery
}

// Handler : charlie rest http handler
type Handler interface {
	Mount(group *echo.Group)
}

// NewHandler : new charlie rest http handler
func NewHandler() Handler {
	return handler{
		deliveryCharlie: deliveryCharlie.NewDelivery(),
	}
}

func (impl handler) Mount(group *echo.Group) {
	group.GET("", impl.getCharlie)
	group.POST("", impl.addCharlie)
	group.PUT("", impl.updateCharlie)
	group.DELETE("", impl.deleteCharlie)
}
