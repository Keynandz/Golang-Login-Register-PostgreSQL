package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Home adalah handler untuk halaman utama.
func Home(c echo.Context) error {
	// Mengirimkan pesan "Selamat Anda Berhasil Terhubung!" dengan status OK (200).
	return c.String(http.StatusOK, "Selamat Anda Berhasil Terhubung!")
}