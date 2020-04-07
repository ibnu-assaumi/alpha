package charlie

import (
	"context"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/tracer"
)

func (impl sql) IsExistsCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.Domain) (exists bool, err error) {

	const operationName string = "Repository_SQL_IsExistsCharlie"

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	subQuery := db.Table(domainCharlie.TableName).
		Select("1").
		Where("status_record <> ?", constant.StatusRecordDelete)

	if param.CharlieID > 0 {
		subQuery = subQuery.Where("charlie_id = ?", param.CharlieID)
	}

	if strings.TrimSpace(param.CharlieName) != "" {
		subQuery = subQuery.Where("charlie_name = ?", param.CharlieName)
	}

	err = db.Raw("SELECT EXISTS ?", subQuery.SubQuery()).Row().Scan(&exists)

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		tracer.SetErrorOpentracing(span, "sql_query", err)
		sentry.CaptureException(err)
		return false, err
	}

	return exists, nil
}
