package charlie

import (
	"context"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	"github.com/olivere/elastic/v7"
)

type es struct{}

// ES : charlie repository elasticsearch
type ES interface {
	GetCharlie(ctx context.Context, client *elastic.Client, param domainCharlie.ParamGet) (result []domainCharlie.Domain, totalData int64, err error)
	UpsertCharlie(ctx context.Context, client *elastic.Client, param domainCharlie.Domain) error
}

// NewES : new charlie repository elasticsearch
func NewES() ES {
	return es{}
}
