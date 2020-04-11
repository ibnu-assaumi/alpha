package beta

import (
	"context"

	"github.com/Bhinneka/alpha/api/lib/response"
	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
	betaUseCase "github.com/Bhinneka/alpha/api/service/usecase/v1/beta"
)

type delivery struct {
	betaUseCase betaUseCase.UseCase
}

type Delivery interface {
	AddBeta(ctx context.Context, param domainBeta.ParamAdd) response.Response
}

func NewDelivery() Delivery {
	return delivery{
		betaUseCase: betaUseCase.NewUseCase(),
	}
}
