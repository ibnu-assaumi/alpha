package beta

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/jsonschema"
	"github.com/Bhinneka/alpha/api/lib/response"
	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
)

func (impl delivery) AddBeta(ctx context.Context, param domainBeta.ParamAdd) response.Response {
	const (
		operationName  string = "Delivery_AddBeta"
		schemaFileName string = "v1/beta/schema_add_beta.json"
	)

	span, spanCTX := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	if err := jsonschema.Validate(schemaFileName, param); err != nil {
		tracer.SetErrorOpentracing(span, constant.ErrorTypeValidateParam, err)
		return response.BadRequest(err.Error())
	}

	result := impl.betaUseCase.AddBeta(spanCTX, param)

	return result
}
