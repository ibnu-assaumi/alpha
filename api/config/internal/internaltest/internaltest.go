package internaltest

import (
	"fmt"
	"os"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// LoadEnv : load env file
func LoadEnv() {
	re := regexp.MustCompile(`^(.*api)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		fmt.Println(".env is not loaded properly")

		os.Exit(2)
	}
}

// SetFakeENV : set fake env
func SetFakeENV() {
	os.Setenv("READ_POSTGRES_DB_USER", "")
	os.Setenv("READ_POSTGRES_DB_PASSWORD", "")
	os.Setenv("READ_POSTGRES_DB_HOST", "")
	os.Setenv("READ_POSTGRES_DB_PORT", "")
	os.Setenv("READ_POSTGRES_DB_NAME", "")
	os.Setenv("READ_POSTGRES_DB_SSL_MODE", "")
	os.Setenv("WRITE_POSTGRES_DB_USER", "")
	os.Setenv("WRITE_POSTGRES_DB_PASSWORD", "")
	os.Setenv("WRITE_POSTGRES_DB_HOST", "")
	os.Setenv("WRITE_POSTGRES_DB_PORT", "")
	os.Setenv("WRITE_POSTGRES_DB_NAME", "")
	os.Setenv("WRITE_POSTGRES_DB_SSL_MODE", "")
}

// GetFakeDB : get fake gorm database client
func GetFakeDB() *gorm.DB {
	sql, _, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", sql)
	return gormDB
}
