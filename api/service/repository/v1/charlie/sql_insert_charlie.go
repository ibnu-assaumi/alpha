package charlie

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

// InsertCharlie : insert new record of charlie to sql database
func (impl sql) InsertCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.Domain) (domainCharlie.Domain, error) {

	const operationName string = "Repository_SQL_InsertCharlie"

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	if err := db.Create(&param).Error; err != nil {
		tracer.SetErrorOpentracing(span, "sql_insert", err)
		sentry.CaptureException(err)
		return param, err
	}

	return param, nil
}
