package charlie

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/response"
	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

func (impl handler) getCharlie(c echo.Context) error {
	const operationName string = "Handler_V1_getCharlie"

	ctx := c.Request().Context()
	span, ctx := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	var param domainCharlie.ParamGet
	if err := c.Bind(&param); err != nil {
		tracer.SetErrorOpentracing(span, constant.ErrorTypeBindParam, err)
		return c.JSON(http.StatusBadRequest, response.BadRequest(err))
	}

	res := impl.deliveryCharlie.GetCharlie(ctx, param)

	return c.JSON(res.Code, res)
}
