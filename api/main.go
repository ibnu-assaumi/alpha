package main

import (
	"context"
	"sync"

	"github.com/joho/godotenv"

	"github.com/Bhinneka/alpha/api/config"
	"github.com/Bhinneka/alpha/api/echo"
	"github.com/Bhinneka/alpha/api/migration"
)

var (
	errorLoadENV = godotenv.Load(".env")
	wg           sync.WaitGroup
)

func main() {
	if errorLoadENV != nil {
		panic(errorLoadENV)
	}

	ctx := context.Background()
	config.Init(ctx)

	defer func() {
		config.PostgresRead.Close()
		config.PostgresWrite.Close()
		config.MongoDB.Client().Disconnect(ctx)
	}()

	migration.MigrateSQL()

	wg.Add(1)
	go func() {
		echo.Start()
	}()
	wg.Wait()
}
