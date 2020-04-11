package beta

import (
	"context"

	"github.com/Bhinneka/alpha/api/lib/response"
	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
	repoBeta "github.com/Bhinneka/alpha/api/service/repository/v1/beta"
	repoCharlie "github.com/Bhinneka/alpha/api/service/repository/v1/charlie"
)

type usecase struct {
	repoBetaSQL    repoBeta.SQL
	repoElastic    repoBeta.ES
	repoMongo      repoBeta.MongoDB
	repoCharlieSQL repoCharlie.SQL
}

type UseCase interface {
	AddBeta(ctx context.Context, param domainBeta.ParamAdd) (result response.Response)
}

func NewUseCase() UseCase {
	return usecase{
		repoBetaSQL:    repoBeta.NewSQL(),
		repoMongo:      repoBeta.NewMongoDB(),
		repoElastic:    repoBeta.NewES(),
		repoCharlieSQL: repoCharlie.NewSQL(),
	}
}
