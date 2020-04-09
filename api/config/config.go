package config

import (
	"context"
	"sync"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/jinzhu/gorm"
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
	ElasticSearch *elasticsearch7.Client
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
