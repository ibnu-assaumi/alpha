package charlie

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Bhinneka/alpha/api/config"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/record"
	mockRepo "github.com/Bhinneka/alpha/api/service/usecase/v1/charlie/internal/sql/mocks"
)

func Test_usecase_GetCharlie(t *testing.T) {
	sql, _, _ := sqlmock.New()
	db, _ := gorm.Open("postgres", sql)

	defer db.Close()

	config.PostgresRead = db

	t.Parallel()

	t.Run("POSITIVE_GET_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamGet{
			CharlieID:   0,
			CharlieName: "name",
			Page:        1,
			Limit:       10,
			OrderBy:     "charlieID",
			Descending:  false,
		}

		resultRepo := []domainCharlie.Domain{
			domainCharlie.Domain{
				CharlieID:   1,
				CharlieName: "name",
				EmbeddedStatus: record.EmbeddedStatus{
					UserIn:       1,
					DateIn:       time.Time{},
					StatusRecord: constant.StatusRecordNew,
				},
			},
		}

		repo := &mockRepo.SQL{}
		repo.On("GetTotalDataCharlie", mock.Anything, mock.Anything, mock.Anything).Return(1, nil)
		repo.On("GetDataCharlie", mock.Anything, mock.Anything, mock.Anything).Return(resultRepo, nil)

		uc := usecase{
			repoSQL: repo,
		}

		result := uc.GetCharlie(ctx, param)

		assert.Equal(t, true, result.Success)
		assert.Equal(t, http.StatusOK, result.Code)
	})

	t.Run("NEGATIVE_PARAM_ID_OR_NAME_GET_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamGet{
			CharlieID:   0,
			CharlieName: "",
			Page:        1,
			Limit:       10,
			OrderBy:     "charlieID",
			Descending:  false,
		}
		repo := &mockRepo.SQL{}
		uc := usecase{
			repoSQL: repo,
		}

		result := uc.GetCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusBadRequest, result.Code)
	})

	t.Run("NEGATIVE_PARAM_PAGE_GET_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamGet{
			CharlieID:   1,
			CharlieName: "",
			Page:        0,
			Limit:       10,
			OrderBy:     "charlieID",
			Descending:  false,
		}
		repo := &mockRepo.SQL{}
		uc := usecase{
			repoSQL: repo,
		}

		result := uc.GetCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusBadRequest, result.Code)
	})

	t.Run("NEGATIVE_PARAM_LIMIT_GET_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamGet{
			CharlieID:   1,
			CharlieName: "",
			Page:        0,
			Limit:       101,
			OrderBy:     "charlieID",
			Descending:  false,
		}
		repo := &mockRepo.SQL{}
		uc := usecase{
			repoSQL: repo,
		}

		result := uc.GetCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusBadRequest, result.Code)
	})

	t.Run("NEGATIVE_PARAM_ORDER_BY_GET_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamGet{
			CharlieID:   1,
			CharlieName: "",
			Page:        0,
			Limit:       100,
			OrderBy:     "test",
			Descending:  false,
		}
		repo := &mockRepo.SQL{}
		uc := usecase{
			repoSQL: repo,
		}

		result := uc.GetCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusBadRequest, result.Code)
	})

	t.Run("NEGATIVE_DATABASE_GET_TOTAL_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamGet{
			CharlieID:   0,
			CharlieName: "name",
			Page:        1,
			Limit:       10,
			OrderBy:     "charlieID",
			Descending:  false,
		}

		repo := &mockRepo.SQL{}
		repo.On("GetTotalDataCharlie", mock.Anything, mock.Anything, mock.Anything).Return(0, errors.New("error"))

		uc := usecase{
			repoSQL: repo,
		}

		result := uc.GetCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusInternalServerError, result.Code)
	})

	t.Run("NEGATIVE_GET_TOTAL_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamGet{
			CharlieID:   0,
			CharlieName: "name",
			Page:        1,
			Limit:       10,
			OrderBy:     "charlieID",
			Descending:  false,
		}

		repo := &mockRepo.SQL{}
		repo.On("GetTotalDataCharlie", mock.Anything, mock.Anything, mock.Anything).Return(0, nil)

		uc := usecase{
			repoSQL: repo,
		}

		result := uc.GetCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusNotFound, result.Code)
	})

	t.Run("NEGATIVE_DATABASE_GET_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamGet{
			CharlieID:   0,
			CharlieName: "name",
			Page:        1,
			Limit:       10,
			OrderBy:     "charlieID",
			Descending:  false,
		}

		repo := &mockRepo.SQL{}
		repo.On("GetTotalDataCharlie", mock.Anything, mock.Anything, mock.Anything).Return(1, nil)
		repo.On("GetDataCharlie", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error"))

		uc := usecase{
			repoSQL: repo,
		}

		result := uc.GetCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusInternalServerError, result.Code)
	})
}
