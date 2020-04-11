package config

import (
	"context"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Bhinneka/alpha/api/config/internal/elasticsearch"
	"github.com/Bhinneka/alpha/api/config/internal/mongodb"
	"github.com/Bhinneka/alpha/api/config/internal/opentracing"
	"github.com/Bhinneka/alpha/api/config/internal/postgresql"
	"github.com/Bhinneka/alpha/api/config/internal/sentry"
)

var (
	initOnce sync.Once
	// PostgresRead : postgresql db read client
	PostgresRead *gorm.DB
	// PostgresWrite : postgresql db write client
	PostgresWrite *gorm.DB
	// MongoDB : mongodb database client
	MongoDB *mongo.Database
	// ElasticSearch : elastic search client
	ElasticSearch *elastic.Client
)

// Init : initialize configuration once
func Init(ctx context.Context) {
	initOnce.Do(func() {
		sentry.InitSentry()
		opentracing.InitOpenTracing()
		PostgresRead = postgresql.GetDBRead()
		PostgresWrite = postgresql.GetDBWrite()
		MongoDB = mongodb.GetDB(ctx)
		ElasticSearch = elasticsearch.GetESClient()
	})
}
