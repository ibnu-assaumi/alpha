package charlie

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/record"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

func Test_sql_InsertCharlie(t *testing.T) {
	t.Parallel()
	t.Run("POSITIVE_INSERT_CHARLIE", func(t *testing.T) {
		sql, mock, _ := sqlmock.New()
		db, _ := gorm.Open("postgres", sql)

		defer db.Close()

		timeNow = func() time.Time {
			return time.Time{}
		}
		ctx := context.Background()
		param := domainCharlie.Domain{
			CharlieName: "name",
			EmbeddedStatus: record.EmbeddedStatus{
				UserIn:       1,
				DateIn:       timeNow(),
				StatusRecord: constant.StatusRecordNew,
			},
		}
		repoSQL := NewSQL()

		rows := mock.NewRows([]string{"charlie_id"}).AddRow(1)
		rowHistory := mock.NewRows([]string{"charlie_id_history"}).AddRow(1)

		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT INTO .*").WillReturnRows(rows)
		mock.ExpectCommit()
		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT INTO .*").WithArgs(
			int64(1),
			param.CharlieName,
			param.UserIn,
			param.DateIn,
			param.StatusRecord).WillReturnRows(rowHistory)
		mock.ExpectCommit()

		_, err := repoSQL.InsertCharlie(ctx, db, param)
		assert.NoError(t, err)
	})

	t.Run("NEGATIVE_INSERT_CHARLIE", func(t *testing.T) {
		sql, mock, _ := sqlmock.New()
		db, _ := gorm.Open("postgres", sql)

		defer db.Close()

		timeNow = func() time.Time {
			return time.Time{}
		}
		ctx := context.Background()
		param := domainCharlie.Domain{
			CharlieName: "name",
			EmbeddedStatus: record.EmbeddedStatus{
				UserIn:       1,
				DateIn:       timeNow(),
				StatusRecord: constant.StatusRecordNew,
			},
		}
		repoSQL := NewSQL()

		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT INTO .*").WillReturnError(errors.New("error"))
		mock.ExpectRollback()

		_, err := repoSQL.InsertCharlie(ctx, db, param)
		assert.Error(t, err)
	})

	t.Run("NEGATIVE_INSERT_CHARLIE_HISTORY", func(t *testing.T) {
		sql, mock, _ := sqlmock.New()
		db, _ := gorm.Open("postgres", sql)

		defer db.Close()

		timeNow = func() time.Time {
			return time.Time{}
		}
		ctx := context.Background()
		param := domainCharlie.Domain{
			CharlieName: "name",
			EmbeddedStatus: record.EmbeddedStatus{
				UserIn:       1,
				DateIn:       timeNow(),
				StatusRecord: constant.StatusRecordNew,
			},
		}
		repoSQL := NewSQL()

		rows := mock.NewRows([]string{"charlie_id"}).AddRow(1)

		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT INTO .*").WillReturnRows(rows)
		mock.ExpectCommit()
		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT INTO .*").WithArgs(
			int64(1),
			param.CharlieName,
			param.UserIn, timeNow(),
			param.StatusRecord).WillReturnError(errors.New("error"))
		mock.ExpectRollback()

		_, err := repoSQL.InsertCharlie(ctx, db, param)
		assert.Error(t, err)
	})
}
