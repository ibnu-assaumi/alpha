package charlie

import (
	"context"

	"github.com/Bhinneka/alpha/api/lib/constant"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	"github.com/Bhinneka/alpha/api/lib/tracer"
)

func (impl sql) DeleteCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.Domain) (result domainCharlie.Domain, err error) {
	const operationName string = "Repository_SQL_DeleteCharlie"

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	err = db.Model(&result).
		Where("status_record <> ?", constant.StatusRecordDelete).
		Where("charlie_id = ?", param.CharlieID).
		Updates(param).Error

	if err != nil {
		tracer.SetErrorOpentracing(span, "sql_update", err)
		sentry.CaptureException(err)
		return result, err
	}

	chrlieHistory := domainCharlie.DomainHistory{
		CharlieID:      result.CharlieID,
		CharlieName:    result.CharlieName,
		EmbeddedStatus: result.EmbeddedStatus,
	}

	if err := db.Create(&chrlieHistory).Error; err != nil {
		tracer.SetErrorOpentracing(span, "sql_insert_history", err)
		sentry.CaptureException(err)
		return result, err
	}

	return result, nil
}
