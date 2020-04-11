package beta

import (
	"context"
	"fmt"

	"github.com/Bhinneka/alpha/api/lib/constant"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"

	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
	domainBetaCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/betacharlie"
	"github.com/opentracing/opentracing-go"
)

func (impl sql) InsertBetaCharlie(ctx context.Context, db *gorm.DB, param domainBeta.Domain) error {

	if len(param.Charlie) == 0 {
		return nil
	}

	const operationName string = "Repository_SQL_InsertBeta"

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	// if err := db.Where("beta_id = ?", param.BetaID).Delete(domainBetaCharlie.Domain{}).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
	// 	tracer.SetErrorOpentracing(span, constant.ErrorTypeSQLInsert, err.Error())
	// 	sentry.CaptureException(err)
	// 	return err
	// }

	for i, v := range param.Charlie {
		betaCharlie := domainBetaCharlie.Domain{
			BetaID:    param.BetaID,
			CharlieID: v.CharlieID,
		}
		if err := db.Create(&betaCharlie).Error; err != nil {
			errs := fmt.Errorf("paramAdd[%v] : %s", i, err.Error())
			tracer.SetErrorOpentracing(span, constant.ErrorTypeSQLInsert, errs.Error())
			sentry.CaptureException(err)
			return err
		}
	}

	return nil
}
