package charlie

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/jsonschema"
	"github.com/Bhinneka/alpha/api/lib/response"
	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

func (impl delivery) AddCharlie(ctx context.Context, param domainCharlie.ParamAdd) response.Response {
	const (
		operationName  string = "Delivery_AddCharlie"
		schemaFileName string = "v1/charlie/schema_add_charlie.json"
	)

	span, spanCTX := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	if err := jsonschema.Validate(schemaFileName, param); err != nil {
		tracer.SetErrorOpentracing(span, constant.ErrorTypeValidateParam, err)
		return response.BadRequest(err.Error())
	}

	result := impl.charlieUseCase.AddCharlie(spanCTX, param)

	return result
}
