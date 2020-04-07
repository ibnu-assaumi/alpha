package postgresql

import (
	"testing"

	"github.com/Bhinneka/alpha/api/config/internal/postgresql/postgresqltest"
)

func Test_setDBCallback(t *testing.T) {
	db := postgresqltest.GetFakeDB()
	defer db.Close()

	t.Run("POSITIVE_SET_DB_CALLBACK", func(t *testing.T) {
		setDBCallback(db)
	})
}

func Test_reportToSentry(t *testing.T) {
	db := postgresqltest.GetFakeDB()
	defer db.Close()

	t.Run("POSITIVE_REPORT_TO_SENTRY", func(t *testing.T) {
		reportToSentry(db.NewScope("test"))
	})
}
