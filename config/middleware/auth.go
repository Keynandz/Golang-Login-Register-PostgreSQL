package middleware

import (
	"net/http"

	"part-1/config/session"
	"github.com/labstack/echo/v4"
)

// RequireLogin adalah middleware yang memastikan pengguna sudah login sebelum mengakses halaman tertentu.
func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Cek apakah sesi pengguna tersedia atau tidak
		if session.GetSession(c, "user") == "" {
			// Jika tidak, redirect ke halaman login
			return c.Redirect(http.StatusFound, "/login")
		}

		// Jika sudah login, lanjutkan eksekusi ke handler berikutnya
		return next(c)
	}
}

// NoCache adalah middleware yang menonaktifkan caching untuk halaman yang dihasilkan.
func NoCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Set header Cache-Control untuk no-cache, no-store, dan must-revalidate
		c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		// Set header Pragma ke no-cache
		c.Response().Header().Set("Pragma", "no-cache")
		// Set header Expires ke "0" untuk memastikan halaman tidak disimpan di cache
		c.Response().Header().Set("Expires", "0")
		// Lanjutkan eksekusi ke handler berikutnya
		return next(c)
	}
}