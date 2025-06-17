package middlewares

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func AuditLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		log.Printf("[AUDIT] %s %s - %s in %v", c.Request().Method, c.Path(), c.RealIP(), time.Since(start))
		return err
	}
}