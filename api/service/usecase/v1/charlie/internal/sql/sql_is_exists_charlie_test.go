package sql

import (
	"context"
	"errors"
	"testing"

	"github.com/Bhinneka/alpha/api/lib/constant"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

func Test_sql_IsExistsCharlie(t *testing.T) {
	t.Parallel()
	t.Run("POSITIVE_IS_EXISTS_CHARLIE", func(t *testing.T) {
		sql, mock, _ := sqlmock.New()
		db, _ := gorm.Open("postgres", sql)

		ctx := context.Background()
		param := domainCharlie.Domain{
			CharlieID:   1,
			CharlieName: "name",
		}
		repoSQL := NewSQL()
		rows := mock.NewRows([]string{"exists"}).AddRow(true)

		mock.ExpectQuery("^SELECT EXISTS .*").WithArgs(constant.StatusRecordDelete, param.CharlieID, param.CharlieName).WillReturnRows(rows)
		_, err := repoSQL.IsExistsCharlie(ctx, db, param)
		assert.NoError(t, err)
	})

	t.Run("NEGATIVE_IS_EXISTS_CHARLIE", func(t *testing.T) {
		sql, mock, _ := sqlmock.New()
		db, _ := gorm.Open("postgres", sql)

		ctx := context.Background()
		param := domainCharlie.Domain{
			CharlieID:   1,
			CharlieName: "name",
		}
		repoSQL := NewSQL()

		mock.ExpectQuery("^SELECT EXISTS .*").WithArgs(constant.StatusRecordDelete, param.CharlieID, param.CharlieName).WillReturnError(errors.New("error"))
		_, err := repoSQL.IsExistsCharlie(ctx, db, param)
		assert.Error(t, err)
	})
}
