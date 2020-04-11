package charlie

import (
	"context"
	"fmt"
	"time"

	"github.com/Bhinneka/alpha/api/lib/constant"

	"github.com/Bhinneka/alpha/api/lib/jsonschema"
	"github.com/Bhinneka/alpha/api/lib/tracer"

	"github.com/Bhinneka/alpha/api/config"
	"github.com/Bhinneka/alpha/api/lib/record"
	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/lib/response"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

// DeleteCharlie : delete existing charlie busines logic process
func (impl usecase) DeleteCharlie(ctx context.Context, param domainCharlie.ParamDelete) (result response.Response) {
	const (
		operationName  string = "UseCase_DeleteCharlie"
		schemaFileName string = "v1/charlie/schema_delete_charlie.json"
	)

	span, spanCTX := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	var err error
	db := config.PostgresWrite.Begin()
	defer record.FinishTransaction(db, err)

	if err := jsonschema.Validate(schemaFileName, param); err != nil {
		tracer.SetErrorOpentracing(span, constant.ErrorTypeValidateParam, err)
		return response.BadRequest(err)
	}

	paramCharlie := domainCharlie.Domain{
		CharlieID: param.CharlieID,
	}

	isExistsCharlie, err := impl.repoSQL.IsExistsCharlie(spanCTX, db, paramCharlie)
	if err != nil {
		return response.InternalServerError()
	}

	if !isExistsCharlie {
		err = fmt.Errorf("charlieID '%v' does not exists", param.CharlieID)
		tracer.SetErrorOpentracing(span, constant.ErrorTypeNotExists, err)
		return response.BadRequest(err)
	}

	paramCharlie.EmbeddedStatus = record.EmbeddedStatus{
		UserUp:       1,
		DateUp:       time.Now().UTC(),
		StatusRecord: constant.StatusRecordDelete,
	}

	data, err := impl.repoSQL.DeleteCharlie(spanCTX, db, paramCharlie)
	if err != nil {
		return response.InternalServerError()
	}

	meta := response.Meta{
		TotalData: 1,
	}

	return response.Success(meta, []domainCharlie.Domain{data}, "success delete charlie")
}
