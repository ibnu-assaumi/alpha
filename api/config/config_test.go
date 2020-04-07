package config

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/joho/godotenv"
)

var _ = func() (_ struct{}) {
	re := regexp.MustCompile(`^(.*api)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		fmt.Println(".env is not loaded properly")

		os.Exit(2)
	}
	return
}()

func TestInit(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		ctx := context.Background()
		Init(ctx)
	})
}
