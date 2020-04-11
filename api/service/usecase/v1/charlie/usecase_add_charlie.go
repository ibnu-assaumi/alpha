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

	if err := jsonschema.Validate(schemaFileName, param); err != nil {
		tracer.SetErrorOpentracing(span, constant.ErrorTypeValidateParam, err)
		return response.BadRequest(err.Error())
	}

	paramCharlie := domainCharlie.Domain{
		CharlieName: param.CharlieName,
	}

	sqlRead := config.PostgresRead

	isCharlieExists, err := impl.repoSQL.IsExistsCharlie(spanCTX, sqlRead, paramCharlie)
	if err != nil {
		return response.InternalServerError()
	}

	if isCharlieExists {
		err = fmt.Errorf("charlieName '%s' already exists", param.CharlieName)
		tracer.SetErrorOpentracing(span, constant.ErrorTypeExists, err)
		return response.BadRequest(err.Error())
	}

	paramCharlie.EmbeddedStatus = record.EmbeddedStatus{
		UserIn:       1,
		DateIn:       time.Now().UTC(),
		StatusRecord: constant.StatusRecordNew,
	}

	sqlWrite := config.PostgresWrite.Begin()
	data, err := impl.repoSQL.InsertCharlie(spanCTX, sqlWrite, paramCharlie)
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

	return response.Success(meta, []domainCharlie.Domain{data}, "success add new charlie")
}
