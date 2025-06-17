package middlewares

import (
	"github.com/labstack/echo/v4"
)

var registry = map[string]echo.MiddlewareFunc{
	"Auth":     Auth,
	"Throttle": Throttle,
	"AuditLog": AuditLog,
}

func Get(name string) echo.MiddlewareFunc {
	return registry[name]
}

func GetWithParam(name, param string) echo.MiddlewareFunc {
	switch name {
	case "RequireRole":
		return RequireRole(param)
	default:
		return nil
	}
}
