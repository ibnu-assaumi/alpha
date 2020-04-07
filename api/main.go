package main

import (
	"sync"

	"github.com/joho/godotenv"

	"github.com/Bhinneka/alpha/api/config"
	"github.com/Bhinneka/alpha/api/echo"
	"github.com/Bhinneka/alpha/api/migration"
)

var (
	errorLoadENV = godotenv.Load(".env")
)

func main() {
	if errorLoadENV != nil {
		panic(errorLoadENV)
	}

	config.Init()

	migration.MigrateSQL()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		echo.Start()
	}()
	wg.Wait()
}
