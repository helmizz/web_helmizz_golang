package middlewares

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
)

// Auth is a middleware that checks if the user is authenticated
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Skip authentication for login page and public paths
		path := c.Path()
		if path == "/" || path == "/login" || strings.HasPrefix(path, "/static") {
			return next(c)
		}
		
		// Check if user is authenticated
		// This is a simplified example; in a real application,
		// you would check session cookies, JWT tokens, etc.
		// token := c.Request().Header.Get("Authorization")
		
		/*
		token :="123456789"
		if token == "" {
			token = c.QueryParam("token")
		}
		
		if token == "" {
			// If there's no authentication header and we're not on a public path,
			// redirect to login
			return c.Redirect(302, "/login")
		}
		*/
		// Here you would validate the token
		// For this example, we'll just log it and proceed
		
		// Ambil session
		sess, err := session.Get("session", c)
		if err != nil {
			return c.Redirect(302, "/")
		}

		// Ambil token dari session
		token, ok := sess.Values["token"].(string)
		if !ok || token == "" {
			return c.Redirect(302, "/")
		}
		
		
		log.Printf("Auth check passed for token: %s", token)
		
		return next(c)
	}
}
