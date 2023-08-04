package middleware

import (
	"net/http"

	"part-1/config/session"
	"github.com/labstack/echo/v4"
)

func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if session.GetSession(c, "user") == "" {
			
			return c.Redirect(http.StatusFound, "/login")
		}

		return next(c)
	}
}

func NoCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Response().Header().Set("Pragma", "no-cache")
		c.Response().Header().Set("expires", "0")
		return next(c)
	}
}