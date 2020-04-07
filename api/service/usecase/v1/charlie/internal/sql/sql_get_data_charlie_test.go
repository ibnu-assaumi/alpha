package sql

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

func Test_sql_GetDataCharlie(t *testing.T) {

	t.Parallel()
	t.Run("POSITIVE_GET_DATA_CHARLIE", func(t *testing.T) {
		sql, mock, _ := sqlmock.New()
		db, _ := gorm.Open("postgres", sql)

		ctx := context.Background()
		param := domainCharlie.ParamGet{
			CharlieID:   1,
			CharlieName: "name",
		}
		repoSQL := NewSQL()
		rows := mock.NewRows([]string{"charlie_id", "charlie_name"}).
			AddRow(1, "one").
			AddRow(2, "two")

		mock.ExpectQuery("^SELECT .*").WillReturnRows(rows)
		_, err := repoSQL.GetDataCharlie(ctx, db, param)
		assert.NoError(t, err)
	})

	t.Run("NEGATIVE_GET_DATA_CHARLIE", func(t *testing.T) {
		sql, mock, _ := sqlmock.New()
		db, _ := gorm.Open("postgres", sql)

		ctx := context.Background()
		param := domainCharlie.ParamGet{
			CharlieID:   1,
			CharlieName: "name",
		}
		repoSQL := NewSQL()

		mock.ExpectQuery("^SELECT .*").WillReturnError(errors.New("error"))
		_, err := repoSQL.GetDataCharlie(ctx, db, param)
		assert.Error(t, err)
	})
}
