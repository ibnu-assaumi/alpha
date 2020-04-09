package charlie

import (
	"context"

	"github.com/Bhinneka/alpha/api/lib/constant"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

// UpdateCharlie : update exsisting charlie data by given parameter from sql database
func (impl sql) UpdateCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.Domain) (result domainCharlie.Domain, err error) {
	const operationName string = "Repository_SQL_UpdateCharlie"

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	err = db.Model(&result).
		Where("status_record <> ?", constant.StatusRecordDelete).
		Updates(param).Error

	if err != nil {
		tracer.SetErrorOpentracing(span, "sql_update", err)
		sentry.CaptureException(err)
		return param, err
	}

	err = db.Find(&result).Error
	if err != nil {
		tracer.SetErrorOpentracing(span, "sql_update", err)
		sentry.CaptureException(err)
		return param, err
	}

	return result, nil
}
