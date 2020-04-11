package charlie

import (
	"context"

	"github.com/Bhinneka/alpha/api/lib/response"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	usecaseCharlie "github.com/Bhinneka/alpha/api/service/usecase/v1/charlie"
)

type delivery struct {
	charlieUseCase usecaseCharlie.UseCase
}

// Delivery : charlie delivery / controller layer to handle parameters and deliver response
type Delivery interface {
	GetCharlie(ctx context.Context, param domainCharlie.ParamGet) response.Response
	AddCharlie(ctx context.Context, param domainCharlie.ParamAdd) response.Response
	UpdateCharlie(ctx context.Context, param domainCharlie.ParamUpdate) response.Response
	DeleteCharlie(ctx context.Context, param domainCharlie.ParamDelete) response.Response
}

// NewDelivery : new charlie delivery / controller layer to handle parameters and deliver response
func NewDelivery() Delivery {
	return delivery{
		charlieUseCase: usecaseCharlie.NewUseCase(),
	}
}
