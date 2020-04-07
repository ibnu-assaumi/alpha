package sql

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Bhinneka/alpha/api/lib/record"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_sql_DeleteCharlie(t *testing.T) {

	sql, mock, _ := sqlmock.New()
	db, _ := gorm.Open("postgres", sql)

	defer db.Close()

	t.Parallel()
	t.Run("POSITIVE_DELETE_CHARLIE", func(t *testing.T) {

		timeNow = func() time.Time {
			return time.Time{}
		}
		ctx := context.Background()
		param := domainCharlie.Domain{
			CharlieID:   1,
			CharlieName: "name",
			EmbeddedStatus: record.EmbeddedStatus{
				UserIn:       1,
				DateIn:       timeNow(),
				UserUp:       1,
				DateUp:       timeNow(),
				StatusRecord: constant.StatusRecordDelete,
			},
		}
		repoSQL := NewSQL()

		rows := mock.NewRows([]string{"charlie_id_history"}).AddRow(1)

		mock.ExpectBegin()
		mock.ExpectExec("^UPDATE .*").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		mock.ExpectBegin()
		// date up is not expected because zero time is considered as null and wont be inserted if sql field allowed null value
		mock.ExpectQuery("^INSERT INTO .*").WithArgs(
			param.CharlieID,
			param.CharlieName,
			param.UserIn,
			param.UserUp,
			param.DateIn,
			constant.StatusRecordDelete).WillReturnRows(rows)
		mock.ExpectCommit()
		_, err := repoSQL.DeleteCharlie(ctx, db, param)
		assert.NoError(t, err)
	})

	t.Run("NEGATIVE_DELETE_CHARLIE", func(t *testing.T) {

		timeNow = func() time.Time {
			return time.Time{}
		}
		ctx := context.Background()
		param := domainCharlie.Domain{
			CharlieID:   1,
			CharlieName: "name",
			EmbeddedStatus: record.EmbeddedStatus{
				UserIn:       1,
				DateIn:       timeNow(),
				UserUp:       1,
				DateUp:       timeNow(),
				StatusRecord: constant.StatusRecordDelete,
			},
		}
		repoSQL := NewSQL()

		mock.ExpectBegin()
		mock.ExpectExec("^UPDATE .*").WillReturnError(errors.New("error"))
		mock.ExpectRollback()
		_, err := repoSQL.DeleteCharlie(ctx, db, param)
		assert.Error(t, err)
	})

	t.Run("NEGATIVE_INSERT_CHARLIE_HISTORY", func(t *testing.T) {

		timeNow = func() time.Time {
			return time.Time{}
		}
		ctx := context.Background()
		param := domainCharlie.Domain{
			CharlieID:   1,
			CharlieName: "name",
			EmbeddedStatus: record.EmbeddedStatus{
				UserIn:       1,
				DateIn:       timeNow(),
				UserUp:       1,
				DateUp:       timeNow(),
				StatusRecord: constant.StatusRecordDelete,
			},
		}
		repoSQL := NewSQL()

		mock.ExpectBegin()
		mock.ExpectExec("^UPDATE .*").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		mock.ExpectBegin()
		// date up is not expected because zero time is considered as null and wont be inserted if sql field allowed null value
		mock.ExpectQuery("^INSERT INTO .*").WithArgs(
			param.CharlieID,
			param.CharlieName,
			param.UserIn,
			param.UserUp,
			param.DateIn,
			constant.StatusRecordDelete).WillReturnError(errors.New("error"))
		mock.ExpectCommit()
		_, err := repoSQL.DeleteCharlie(ctx, db, param)
		assert.Error(t, err)
	})
}
