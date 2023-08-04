package handlers

import (
	"fmt"
	"net/http"

	"part-1/config/session"

	"github.com/labstack/echo/v4"
)

// HandleLogout adalah fungsi untuk menangani request dari endpoint logout.
func HandleLogout(c echo.Context) error {
	// Membersihkan session pengguna.
	session.ClearSession(c)

	// Menampilkan pesan ke konsol bahwa session telah dibersihkan.
	fmt.Println("Menghapus session")

	// Redirect pengguna ke halaman "/logout" setelah logout berhasil.
	return c.Redirect(http.StatusFound, "/logout")
}