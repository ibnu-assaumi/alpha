package beta

import (
	"github.com/labstack/echo/v4"

	deliveryBeta "github.com/Bhinneka/alpha/api/service/delivery/v1/beta"
)

type handler struct {
	deliveryBeta deliveryBeta.Delivery
}

// Handler : beta rest http handler
type Handler interface {
	Mount(group *echo.Group)
}

// NewHandler : new beta rest http handler
func NewHandler() Handler {
	return handler{
		deliveryBeta: deliveryBeta.NewDelivery(),
	}
}

func (impl handler) Mount(group *echo.Group) {
	group.POST("", impl.addBeta)
}
