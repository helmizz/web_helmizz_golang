package middlewares

import "github.com/labstack/echo/v4"

func Chain(h echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) echo.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
