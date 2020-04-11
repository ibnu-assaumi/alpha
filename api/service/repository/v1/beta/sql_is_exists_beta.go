package beta

import (
	"context"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
)

func (impl sql) IsExistsBeta(ctx context.Context, db *gorm.DB, param domainBeta.Domain) (exists bool, err error) {

	const operationName string = "Repository_SQL_IsExistsBeta"

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	subQuery := db.Table(domainBeta.TableName).
		Select("1").
		Where("status_record <> ?", constant.StatusRecordDelete)

	if param.BetaID > 0 {
		subQuery = subQuery.Where("beta_id = ?", param.BetaID)
	}

	if strings.TrimSpace(param.BetaName) != "" {
		subQuery = subQuery.Where("beta_name = ?", param.BetaName)
	}

	err = db.Raw("SELECT EXISTS ?", subQuery.SubQuery()).Row().Scan(&exists)

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		tracer.SetErrorOpentracing(span, constant.ErrorTypeSQLQuery, err)
		sentry.CaptureException(err)
		return false, err
	}

	return exists, nil
}
