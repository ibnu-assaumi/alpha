package postgresql

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

func setDBStats(db *gorm.DB, dbType string) {
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	if dbType != dbTypeRead && dbType != dbTypeWrite {
		panic(errors.New("dbType must be 'READ' or 'WRITE'"))
	}

	logMode, err := strconv.ParseBool(os.Getenv(fmt.Sprintf("%s_POSTGRES_DB_LOG_MODE", dbType)))
	if err != nil {
		logMode = defaultLogMode
	}
	db.LogMode(logMode)

	maxIdleConnection, err := strconv.Atoi(os.Getenv(fmt.Sprintf("%s_POSTGRES_DB_MAX_IDLE_CONN", dbType)))
	if err != nil {
		maxIdleConnection = defaultMaxIdleConnection
	}
	db.DB().SetMaxIdleConns(maxIdleConnection)

	maxOpenConnection, err := strconv.Atoi(os.Getenv(fmt.Sprintf("%s_POSTGRES_DB_MAX_OPEN_CONN", dbType)))
	if err != nil {
		maxOpenConnection = defaultMaxOpenConnection
	}

	db.DB().SetMaxOpenConns(maxOpenConnection)

	maxLifeTime, err := strconv.Atoi(os.Getenv(fmt.Sprintf("%s_POSTGRES_DB_MAX_LIFETIME", dbType)))
	if err != nil {
		maxLifeTime = defaultMaxLifeTime
	}
	connMaxLifeTime := time.Duration(maxLifeTime) * time.Second

	db.DB().SetConnMaxLifetime(connMaxLifeTime)
}
