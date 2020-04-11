package beta

import (
	"go.mongodb.org/mongo-driver/mongo"

	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
)

type mongoDB struct{}

type MongoDB interface {
	InsertBetaHistory(db *mongo.Database, param domainBeta.Domain)
}

func NewMongoDB() MongoDB {
	return mongoDB{}
}
