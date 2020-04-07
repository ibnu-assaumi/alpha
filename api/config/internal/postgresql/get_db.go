package postgresql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres dialect
)

const (
	dbTypeRead               string = "READ"
	dbTypeWrite              string = "WRITE"
	defaultLogMode           bool   = false
	defaultMaxOpenConnection int    = 10
	defaultMaxIdleConnection int    = 5
	defaultMaxLifeTime       int    = 10
)

// GetDBRead : get postgresql database read connection
func GetDBRead() *gorm.DB {
	connectionString := getConnectionString(dbTypeRead)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	setDBCallback(db)
	setDBStats(db, dbTypeRead)

	return db

}

// GetDBWrite : get postgresql database write connection
func GetDBWrite() *gorm.DB {
	connectionString := getConnectionString(dbTypeWrite)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	setDBCallback(db)
	setDBStats(db, dbTypeWrite)

	return db

}
