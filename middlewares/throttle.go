package middlewares

import (
	"net"
	"time"

	"github.com/labstack/echo/v4"
)

var ipMap = map[string]time.Time{}

func Throttle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip, _, _ := net.SplitHostPort(c.Request().RemoteAddr)
		lastAccess := ipMap[ip]
		if time.Since(lastAccess) < 2*time.Second {
			return c.String(429, "Too many requests")
		}
		ipMap[ip] = time.Now()
		return next(c)
	}
}
