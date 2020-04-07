package charlie

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/config"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/jsonschema"
	"github.com/Bhinneka/alpha/api/lib/record"
	"github.com/Bhinneka/alpha/api/lib/response"
	"github.com/Bhinneka/alpha/api/lib/tracer"
)

// UpdateCharlie : update existing charlie business logic process
func (impl usecase) UpdateCharlie(ctx context.Context, param domainCharlie.ParamUpdate) (result response.Response) {
	const (
		operationName  string = "UseCase_UpdateCharlie"
		schemaFileName string = "v1/charlie/schema_update_charlie.json"
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
		CharlieID: param.CharlieID,
	}

	isCharlieExists, err := impl.repoSQL.IsExistsCharlie(spanCTX, db, paramCharlie)
	if err != nil {
		return response.InternalServerError()
	}

	if !isCharlieExists {
		err = fmt.Errorf("charlieID '%v' does not exists", param.CharlieID)
		tracer.SetErrorOpentracing(span, constant.ErrorTypeNotExists, err)
		return response.BadRequest(err.Error())
	}

	paramCharlie.CharlieName = param.CharlieName

	embeddedStatus := record.EmbeddedStatus{
		UserUp:       1,
		DateUp:       time.Now(),
		StatusRecord: constant.StatusRecordUpdate,
	}

	paramCharlie.EmbeddedStatus = embeddedStatus

	data, err := impl.repoSQL.UpdateCharlie(spanCTX, db, paramCharlie)
	if err != nil {
		return response.InternalServerError()
	}

	meta := response.Meta{
		TotalData: 1,
	}

	return response.Success(meta, []domainCharlie.Domain{data}, "success update charlie")
}
