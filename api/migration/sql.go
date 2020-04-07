package migration

import (
	"errors"
	"sync"

	"github.com/Bhinneka/alpha/api/config"

	"github.com/Bhinneka/alpha/api/migration/sql"
)

// MigrateSQL : migrate sql tables
func MigrateSQL() {
	db := config.PostgresWrite

	if db == nil {
		panic(errors.New("database connection is nil"))
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		sql.MigrateCharlie(db)
		wg.Done()
	}()

	wg.Wait()
}
