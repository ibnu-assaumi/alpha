package config

import (
	"sync"

	"github.com/jinzhu/gorm"

	"github.com/Bhinneka/alpha/api/config/internal/opentracing"
	"github.com/Bhinneka/alpha/api/config/internal/postgresql"
	"github.com/Bhinneka/alpha/api/config/internal/sentry"
)

var (
	initOnce sync.Once
	// PostgresRead : postgresql db read
	PostgresRead *gorm.DB
	// PostgresWrite : postgresql db write
	PostgresWrite *gorm.DB
)

// Init : initialize configuration once
func Init() {
	initOnce.Do(func() {
		sentry.InitSentry()
		opentracing.InitOpenTracing()
		PostgresRead = postgresql.GetDBRead()
		PostgresWrite = postgresql.GetDBWrite()
	})
}
