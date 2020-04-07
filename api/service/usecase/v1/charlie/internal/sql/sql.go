package sql

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

var timeNow = time.Now

type sql struct{}

// SQL : charlie sql repository
type SQL interface {
	GetDataCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.ParamGet) ([]domainCharlie.Domain, error)
	GetTotalDataCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.ParamGet) (int, error)
	IsExistsCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.Domain) (bool, error)
	InsertCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.Domain) (domainCharlie.Domain, error)
	UpdateCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.Domain) (domainCharlie.Domain, error)
	DeleteCharlie(ctx context.Context, db *gorm.DB, param domainCharlie.Domain) (domainCharlie.Domain, error)
}

// NewSQL : new charlie sql repository
func NewSQL() SQL {
	return sql{}
}
