package beta

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/config"
	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/record"
	"github.com/Bhinneka/alpha/api/lib/response"
	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

// AddBeta : add beta business logic process
func (impl usecase) AddBeta(ctx context.Context, param domainBeta.ParamAdd) (result response.Response) {
	const operationName string = "UseCase_AddBeta"

	span, spanCTX := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	paramBeta := domainBeta.Domain{
		BetaName: param.BetaName,
	}

	sqlRead := config.PostgresRead

	isBetaExists, err := impl.repoBetaSQL.IsExistsBeta(spanCTX, sqlRead, paramBeta)
	if err != nil {
		return response.InternalServerError()
	}

	if isBetaExists {
		err = fmt.Errorf("betaName '%s' already exists", param.BetaName)
		tracer.SetErrorOpentracing(span, constant.ErrorTypeExists, err)
		return response.BadRequest(err.Error())
	}

	charlieList := make([]domainCharlie.Domain, 0)
	for i, v := range param.CharlieIDList {
		paramCharlie := domainCharlie.ParamGet{
			CharlieID: v,
			Page:      1,
			Limit:     1,
		}
		dataCharlie, err := impl.repoCharlieSQL.GetDataCharlie(spanCTX, sqlRead, paramCharlie)
		if err != nil {
			return response.InternalServerError()
		}

		if len(dataCharlie) == 0 {
			err = fmt.Errorf("charlieIDList[%v] : '%v' does not exists", i, v)
			tracer.SetErrorOpentracing(span, constant.ErrorTypeNotExists, err)
			return response.BadRequest(err.Error())
		}

		charlieList = append(charlieList, dataCharlie[0])
	}

	paramBeta.EmbeddedStatus = record.EmbeddedStatus{
		UserIn:       1,
		DateIn:       time.Now().UTC(),
		StatusRecord: constant.StatusRecordNew,
	}

	sqlWrite := config.PostgresWrite.Begin()
	data, err := impl.repoBetaSQL.InsertBeta(spanCTX, sqlWrite, paramBeta)
	if err != nil {
		sqlWrite.Rollback()
		return response.InternalServerError()
	}

	data.Charlie = charlieList

	if err := impl.repoBetaSQL.InsertBetaCharlie(spanCTX, sqlWrite, data); err != nil {
		sqlWrite.Rollback()
		return response.InternalServerError()
	}

	if err := impl.repoElastic.UpsertBeta(spanCTX, config.ElasticSearch, data); err != nil {
		sqlWrite.Rollback()
		return response.InternalServerError()
	}

	sqlWrite.Commit()

	go impl.repoMongo.InsertBetaHistory(config.MongoDB, data)

	meta := response.Meta{
		TotalData: 1,
	}

	return response.Success(meta, []domainBeta.Domain{data}, "success add new beta")
}
