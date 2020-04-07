package charlie

import (
	"context"
	"math"

	"github.com/Bhinneka/alpha/api/lib/constant"

	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/config"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	"github.com/Bhinneka/alpha/api/lib/jsonschema"
	"github.com/Bhinneka/alpha/api/lib/response"
	"github.com/Bhinneka/alpha/api/lib/tracer"
)

// GetCharlie : get list of charlie business logic process
func (impl usecase) GetCharlie(ctx context.Context, param domainCharlie.ParamGet) (result response.Response) {
	const (
		operationName  string = "UseCase_GetCharlie"
		schemaFileName string = "v1/charlie/schema_get_charlie.json"
	)

	span, spanCTX := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	if err := jsonschema.Validate(schemaFileName, param); err != nil {
		tracer.SetErrorOpentracing(span, constant.ErrorTypeValidateParam, err)
		return response.BadRequest(err)
	}

	db := config.PostgresRead
	totalData, err := impl.repoSQL.GetTotalDataCharlie(spanCTX, db, param)
	if err != nil {
		return response.InternalServerError()
	}

	if totalData < 1 {
		errMessage := "record not found"
		tracer.SetErrorOpentracing(span, constant.ErrorTypeNotExists, errMessage)
		return response.NotFound(errMessage)
	}

	data, err := impl.repoSQL.GetDataCharlie(spanCTX, db, param)
	if err != nil {
		return response.InternalServerError()
	}

	pages := float64(totalData / param.Limit)
	totalPage := int(math.Floor(pages)) + 1
	meta := response.Meta{
		Limit:     param.Limit,
		Page:      param.Page,
		TotalData: totalData,
		TotalPage: totalPage,
	}

	return response.Success(meta, data, "success get charlie")
}
