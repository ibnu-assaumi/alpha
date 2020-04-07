package sql

import (
	"context"
	"fmt"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/tracer"
)

// GetTotalDataCharlie : get total data of charlie filtered by given parameter from sql database
func (impl sql) GetTotalDataCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.ParamGet) (count int, err error) {
	const operationName string = "Repository_SQL_GetTotalDataCharlie"

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	if param.CharlieID != 0 {
		db = db.Where("charlie_id = ?", param.CharlieID)
	}

	if strings.TrimSpace(param.CharlieName) != "" {
		db = db.Where("charlie_name LIKE ?", fmt.Sprintf("%%%s%%", param.CharlieName))
	}

	err = db.Table(domainCharlie.TableName).
		Select("charlie_id").
		Where("status_record <> ?", constant.StatusRecordDelete).
		Count(&count).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		tracer.SetErrorOpentracing(span, "sql_query", err)
		sentry.CaptureException(err)
		return count, err
	}

	return count, err
}
