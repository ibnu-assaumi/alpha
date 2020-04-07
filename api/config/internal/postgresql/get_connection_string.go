package postgresql

import (
	"errors"
	"fmt"
	"os"
)

func getConnectionString(dbType string) string {

	if dbType != dbTypeRead && dbType != dbTypeWrite {
		panic(errors.New("dbType must be 'READ' or 'WRITE'"))
	}

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv(fmt.Sprintf("%s_POSTGRES_DB_USER", dbType)),
		os.Getenv(fmt.Sprintf("%s_POSTGRES_DB_PASSWORD", dbType)),
		os.Getenv(fmt.Sprintf("%s_POSTGRES_DB_HOST", dbType)),
		os.Getenv(fmt.Sprintf("%s_POSTGRES_DB_PORT", dbType)),
		os.Getenv(fmt.Sprintf("%s_POSTGRES_DB_NAME", dbType)),
		os.Getenv(fmt.Sprintf("%s_POSTGRES_DB_SSL_MODE", dbType)),
	)
}
