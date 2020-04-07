package charlie

import (
	"context"

	"github.com/Bhinneka/alpha/api/lib/response"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	sqlCharlie "github.com/Bhinneka/alpha/api/service/repository/v1/charlie"
)

type usecase struct {
	repoSQL sqlCharlie.SQL
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
		repoSQL: sqlCharlie.NewSQL(),
	}
}
