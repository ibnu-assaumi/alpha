package sql

import (
	"context"
	"fmt"
	"strings"

	"github.com/Bhinneka/alpha/api/lib/stringutil"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/tracer"
)

// GetDataCharlie : get list of charlie data filtered by given parameter from sql database
func (impl sql) GetDataCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.ParamGet) (result []domainCharlie.Domain, err error) {
	const operationName string = "Repository_SQL_GetTotalDataCharlie"

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	offset := (param.Page - 1) * param.Limit

	orderBy := stringutil.CamelToSnakeCase(param.OrderBy)
	if orderBy == "" {
		orderBy = "charlie_id"
	}

	if param.Descending {
		orderBy += " desc"
	} else {
		orderBy += " asc"
	}

	if param.CharlieID != 0 {
		db = db.Where("charlie_id = ?", param.CharlieID)
	}

	if strings.TrimSpace(param.CharlieName) != "" {
		db = db.Where("charlie_name LIKE ?", fmt.Sprintf("%%%s%%", param.CharlieName))
	}

	err = db.Where("status_record <> ?", constant.StatusRecordDelete).
		Offset(offset).
		Limit(param.Limit).
		Order(orderBy).
		Find(&result).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		tracer.SetErrorOpentracing(span, "sql_query", err)
		sentry.CaptureException(err)
		return result, err
	}

	return result, err
}
