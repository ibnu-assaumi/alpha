package charlie

import (
	"context"

	"github.com/Bhinneka/alpha/api/lib/response"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	repoCharlie "github.com/Bhinneka/alpha/api/service/repository/v1/charlie"
)

type usecase struct {
	repoSQL     repoCharlie.SQL
	repoMongo   repoCharlie.MongoDB
	repoElastic repoCharlie.ES
}

// UseCase : charlie business logic layer
type UseCase interface {
	GetCharlie(ctx context.Context, param domainCharlie.ParamGet) (result response.Response)
	AddCharlie(ctx context.Context, param domainCharlie.ParamAdd) (result response.Response)
	UpdateCharlie(ctx context.Context, param domainCharlie.ParamUpdate) (result response.Response)
	DeleteCharlie(ctx context.Context, param domainCharlie.ParamDelete) (result response.Response)
}

// NewUseCase : new charlie business logic layer
func NewUseCase() UseCase {
	return usecase{
		repoSQL:     repoCharlie.NewSQL(),
		repoMongo:   repoCharlie.NewMongoDB(),
		repoElastic: repoCharlie.NewES(),
	}
}
