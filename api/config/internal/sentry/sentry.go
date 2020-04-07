package sentry

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
)

// InitSentry : init sentry once
func InitSentry() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		panic(err)
	}
}
