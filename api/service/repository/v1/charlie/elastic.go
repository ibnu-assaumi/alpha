package charlie

import (
	"context"

	"github.com/elastic/go-elasticsearch/v7"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

type elastic struct{}

// Elastic : charlie repository elastic
type Elastic interface {
	UpsertCharlie(ctx context.Context, db *elasticsearch.Client, param domainCharlie.Domain)
}

// NewElastic : new charlie repository elastic
func NewElastic() Elastic {
	return elastic{}
}
