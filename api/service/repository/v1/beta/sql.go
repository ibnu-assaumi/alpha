package beta

import (
	"context"

	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
	"github.com/jinzhu/gorm"
)

type sql struct{}

type SQL interface {
	InsertBeta(ctx context.Context, db *gorm.DB, param domainBeta.Domain) (result domainBeta.Domain, err error)
	InsertBetaCharlie(ctx context.Context, db *gorm.DB, param domainBeta.Domain) error
	IsExistsBeta(ctx context.Context, db *gorm.DB, param domainBeta.Domain) (exists bool, err error)
}

func NewSQL() SQL {
	return sql{}
}
