package beta

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"

	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
	"github.com/opentracing/opentracing-go"
)

func (impl sql) InsertBeta(ctx context.Context, db *gorm.DB, param domainBeta.Domain) (result domainBeta.Domain, err error) {
	const operationName string = "Repository_SQL_InsertBeta"

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()
	if err := db.Create(&param).Error; err != nil {
		tracer.SetErrorOpentracing(span, constant.ErrorTypeSQLInsert, err)
		sentry.CaptureException(err)
		return param, err
	}

	return param, nil
}
