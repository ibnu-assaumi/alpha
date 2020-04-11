package charlie

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/config"
	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/jsonschema"
	"github.com/Bhinneka/alpha/api/lib/record"
	"github.com/Bhinneka/alpha/api/lib/response"
	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

// UpdateCharlie : update existing charlie business logic process
func (impl usecase) UpdateCharlie(ctx context.Context, param domainCharlie.ParamUpdate) (result response.Response) {
	const (
		operationName  string = "UseCase_UpdateCharlie"
		schemaFileName string = "v1/charlie/schema_update_charlie.json"
	)

	span, spanCTX := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	if err := jsonschema.Validate(schemaFileName, param); err != nil {
		tracer.SetErrorOpentracing(span, constant.ErrorTypeValidateParam, err)
		return response.BadRequest(err.Error())
	}

	paramCharlie := domainCharlie.Domain{
		CharlieID: param.CharlieID,
	}

	sqlWrite := config.PostgresWrite.Begin()
	isCharlieExists, err := impl.repoSQL.IsExistsCharlie(spanCTX, sqlWrite, paramCharlie)
	if err != nil {
		sqlWrite.Rollback()
		return response.InternalServerError()
	}

	if !isCharlieExists {
		sqlWrite.Rollback()
		err = fmt.Errorf("charlieID '%v' does not exists", param.CharlieID)
		tracer.SetErrorOpentracing(span, constant.ErrorTypeNotExists, err)
		return response.BadRequest(err.Error())
	}

	paramCharlie.CharlieName = param.CharlieName
	paramCharlie.EmbeddedStatus = record.EmbeddedStatus{
		UserUp:       1,
		DateUp:       time.Now().UTC(),
		StatusRecord: constant.StatusRecordUpdate,
	}

	data, err := impl.repoSQL.UpdateCharlie(spanCTX, sqlWrite, paramCharlie)
	if err != nil {
		sqlWrite.Rollback()
		return response.InternalServerError()
	}

	if err := impl.repoElastic.UpsertCharlie(spanCTX, config.ElasticSearch, data); err != nil {
		sqlWrite.Rollback()
		return response.InternalServerError()
	}

	sqlWrite.Commit()

	go impl.repoMongo.InsertCharlieHistory(config.MongoDB, data)

	meta := response.Meta{
		TotalData: 1,
	}

	return response.Success(meta, []domainCharlie.Domain{data}, "success update charlie")
}
