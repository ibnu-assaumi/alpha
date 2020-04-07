package opentracing

import (
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

func TestInitOpenTracing(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		InitOpenTracing()
	})

	t.Run("PANIC", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		os.Setenv("JAEGER_SERVICE_NAME", "")
		InitOpenTracing()
	})
}
