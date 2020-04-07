package charlie

import (
	"context"
	"fmt"
	"time"

	"github.com/Bhinneka/alpha/api/lib/jsonschema"
	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/config"
	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/record"
	"github.com/Bhinneka/alpha/api/lib/response"
	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

// AddCharlie : add charlie business logic process
func (impl usecase) AddCharlie(ctx context.Context, param domainCharlie.ParamAdd) (result response.Response) {
	const (
		operationName  string = "UseCase_AddCharlie"
		schemaFileName string = "v1/charlie/schema_add_charlie.json"
	)

	span, spanCTX := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	var err error
	db := config.PostgresWrite.Begin()
	defer record.FinishTransaction(db, err)

	if err = jsonschema.Validate(schemaFileName, param); err != nil {
		tracer.SetErrorOpentracing(span, constant.ErrorTypeValidateParam, err)
		return response.BadRequest(err.Error())
	}

	paramCharlie := domainCharlie.Domain{
		CharlieName: param.CharlieName,
	}

	isCharlieExists, err := impl.repoSQL.IsExistsCharlie(spanCTX, db, paramCharlie)
	if err != nil {
		return response.InternalServerError()
	}

	if isCharlieExists {
		err = fmt.Errorf("charlieName '%s' already exists", param.CharlieName)
		tracer.SetErrorOpentracing(span, constant.ErrorTypeExists, err)
		return response.BadRequest(err.Error())
	}

	status := record.EmbeddedStatus{
		UserIn:       1,
		DateIn:       time.Now(),
		StatusRecord: constant.StatusRecordNew,
	}

	paramCharlie.EmbeddedStatus = status

	data, err := impl.repoSQL.InsertCharlie(spanCTX, db, paramCharlie)
	if err != nil {
		return response.InternalServerError()
	}

	meta := response.Meta{
		TotalData: 1,
	}

	return response.Success(meta, []domainCharlie.Domain{data}, "success add new charlie")
}