package beta

import (
	"context"

	"github.com/olivere/elastic/v7"

	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
)

type es struct{}

type ES interface {
	UpsertBeta(ctx context.Context, client *elastic.Client, param domainBeta.Domain) error
}

func NewES() ES {
	return es{}
}
