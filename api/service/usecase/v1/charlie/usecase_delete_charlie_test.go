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

func Test_usecase_DeleteCharlie(t *testing.T) {
	sql, mockDB, _ := sqlmock.New()
	db, _ := gorm.Open("postgres", sql)

	defer db.Close()

	config.PostgresWrite = db

	t.Parallel()

	t.Run("POSITIVE_DELETE_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamDelete{
			CharlieID: 1,
		}

		paramRepo := domainCharlie.Domain{
			CharlieID: param.CharlieID,
			EmbeddedStatus: record.EmbeddedStatus{
				UserIn:       1,
				DateIn:       time.Time{},
				UserUp:       1,
				DateUp:       time.Time{},
				StatusRecord: constant.StatusRecordDelete,
			},
		}

		repo := &mockRepo.SQL{}
		repo.On("IsExistsCharlie", mock.Anything, mock.Anything, mock.Anything).Return(true, nil)
		repo.On("DeleteCharlie", mock.Anything, mock.Anything, mock.Anything).Return(paramRepo, nil)

		uc := &usecase{
			repoSQL: repo,
		}

		mockDB.ExpectBegin()
		mockDB.ExpectCommit()

		result := uc.DeleteCharlie(ctx, param)

		assert.Equal(t, true, result.Success)
		assert.Equal(t, http.StatusOK, result.Code)
	})

	t.Run("NEGATIVE_VALIDATION_DELETE_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamDelete{
			CharlieID: 0,
		}

		repo := &mockRepo.SQL{}

		uc := &usecase{
			repoSQL: repo,
		}

		mockDB.ExpectBegin()
		mockDB.ExpectRollback()

		result := uc.DeleteCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusBadRequest, result.Code)
	})

	t.Run("NEGATIVE_IS_EXISTS_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamDelete{
			CharlieID: 1,
		}

		repo := &mockRepo.SQL{}
		repo.On("IsExistsCharlie", mock.Anything, mock.Anything, mock.Anything).Return(false, nil)

		uc := &usecase{
			repoSQL: repo,
		}

		mockDB.ExpectBegin()
		mockDB.ExpectRollback()

		result := uc.DeleteCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusBadRequest, result.Code)
	})

	t.Run("NEGATIVE_DATABASE_IS_EXISTS_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamDelete{
			CharlieID: 1,
		}

		repo := &mockRepo.SQL{}
		repo.On("IsExistsCharlie", mock.Anything, mock.Anything, mock.Anything).Return(false, errors.New("error"))

		uc := &usecase{
			repoSQL: repo,
		}

		mockDB.ExpectBegin()
		mockDB.ExpectRollback()

		result := uc.DeleteCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusInternalServerError, result.Code)
	})

	t.Run("NEGATIVE_DATABASE_DELETE_CHARLIE", func(t *testing.T) {
		ctx := context.Background()
		param := domainCharlie.ParamDelete{
			CharlieID: 1,
		}

		repo := &mockRepo.SQL{}
		repo.On("IsExistsCharlie", mock.Anything, mock.Anything, mock.Anything).Return(true, nil)
		repo.On("DeleteCharlie", mock.Anything, mock.Anything, mock.Anything).Return(domainCharlie.Domain{}, errors.New("error"))

		uc := &usecase{
			repoSQL: repo,
		}

		mockDB.ExpectBegin()
		mockDB.ExpectRollback()

		result := uc.DeleteCharlie(ctx, param)

		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusInternalServerError, result.Code)
	})
}
