package echo

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

// Start : echo start http rest server
func Start() {
	e := echo.New()

	spanStart := opentracing.StartSpan("Start_Echo")
	spanStart.Finish()

	setupMiddleware(e)
	setupHandler(e)

	if err := e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))); err != nil {
		panic(err)
	}
}
