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
	mockRepo "github.com/Bhinneka/alpha/api/service/repository/v1/charlie/mocks"
)

func Test_usecase_UpdateCharlie(t *testing.T) {
	sql, mockDB, _ := sqlmock.New()
	db, _ := gorm.Open("postgres", sql)

	defer db.Close()

	config.PostgresWrite = db

	t.Parallel()

	t.Run("POSITIVE_UPDATE_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamUpdate{
			CharlieID:   1,
			CharlieName: "name",
		}

		paramRepo := domainCharlie.Domain{
			CharlieID:   1,
			CharlieName: param.CharlieName,
			EmbeddedStatus: record.EmbeddedStatus{
				UserUp:       1,
				DateUp:       time.Now(),
				StatusRecord: constant.StatusRecordUpdate,
			},
		}

		repo := &mockRepo.SQL{}
		repo.On("IsExistsCharlie", mock.Anything, mock.Anything, mock.Anything).Return(true, nil)
		repo.On("UpdateCharlie", mock.Anything, mock.Anything, mock.Anything).Return(paramRepo, nil)

		uc := &usecase{
			repoSQL: repo,
		}

		mockDB.ExpectBegin()
		mockDB.ExpectCommit()

		result := uc.UpdateCharlie(ctx, param)

		assert.Equal(t, true, result.Success)
		assert.Equal(t, http.StatusOK, result.Code)
	})

	t.Run("NEGATIVE_VALIDATE_UPDATE_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamUpdate{
			CharlieID:   0,
			CharlieName: "",
		}

		repo := &mockRepo.SQL{}
		uc := &usecase{
			repoSQL: repo,
		}

		mockDB.ExpectBegin()
		mockDB.ExpectRollback()

		result := uc.UpdateCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusBadRequest, result.Code)
	})

	t.Run("NEGATIVE_IS_CHARLIE_EXISTS", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamUpdate{
			CharlieID:   1,
			CharlieName: "name",
		}

		repo := &mockRepo.SQL{}
		repo.On("IsExistsCharlie", mock.Anything, mock.Anything, mock.Anything).Return(false, nil)

		uc := &usecase{
			repoSQL: repo,
		}

		mockDB.ExpectBegin()
		mockDB.ExpectRollback()

		result := uc.UpdateCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusBadRequest, result.Code)
	})

	t.Run("NEGATIVE_DATABASE_IS_CHARLIE_EXISTS", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamUpdate{
			CharlieID:   1,
			CharlieName: "name",
		}

		repo := &mockRepo.SQL{}
		repo.On("IsExistsCharlie", mock.Anything, mock.Anything, mock.Anything).Return(false, errors.New("error"))

		uc := &usecase{
			repoSQL: repo,
		}

		mockDB.ExpectBegin()
		mockDB.ExpectRollback()

		result := uc.UpdateCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusInternalServerError, result.Code)
	})

	t.Run("NEGATIVE_DATABASE_IS_CHARLIE_EXISTS", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamUpdate{
			CharlieID:   1,
			CharlieName: "name",
		}

		repo := &mockRepo.SQL{}
		repo.On("IsExistsCharlie", mock.Anything, mock.Anything, mock.Anything).Return(true, nil)
		repo.On("UpdateCharlie", mock.Anything, mock.Anything, mock.Anything).Return(domainCharlie.Domain{}, errors.New("error"))

		uc := &usecase{
			repoSQL: repo,
		}

		mockDB.ExpectBegin()
		mockDB.ExpectRollback()

		result := uc.UpdateCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusInternalServerError, result.Code)
	})
}
