package middleware

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// TimeoutRequest : middleware request timeout
func TimeoutRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		timeout, err := strconv.Atoi(os.Getenv("SERVER_REQUEST_TIMEOUT_SECOND"))
		if err != nil {
			timeout = 120
		}

		req := c.Request()
		ctx, cancel := context.WithTimeout(req.Context(), time.Duration(timeout)*time.Second)
		defer cancel()

		c.SetRequest(req.WithContext(ctx))
		return next(c)
	}
}
