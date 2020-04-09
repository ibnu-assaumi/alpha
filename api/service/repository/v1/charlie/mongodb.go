package charlie

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

type mongoDB struct{}

// MongoDB : charlie mongodb repository
type MongoDB interface {
	InsertCharlieHistory(ctx context.Context, db *mongo.Database, param domainCharlie.Domain)
}

// NewMongoDB : new charlie mongodb repository
func NewMongoDB() MongoDB {
	return mongoDB{}
}
