package handlers

import (
	"part-1/config/repositories"
	"part-1/config/session"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// HandleLogin adalah fungsi untuk menangani request dari endpoint login.
func HandleLogin(c echo.Context) error {
	contentType := c.Request().Header.Get("Content-Type")
	if contentType == "application/json" {
		return LoginUser(c)
	}
	return LoginUserWeb(c)
}

// LoginUserWeb adalah fungsi untuk melakukan proses login melalui web dengan form values.
func LoginUserWeb(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Cari pengguna berdasarkan email yang diberikan.
	user, err := repositories.GetAkunByEmail(email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Email atau kata sandi salah")
	}

	// Bandingkan kata sandi yang dimasukkan dengan kata sandi yang ada di database.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Email atau kata sandi salah")
	}

	// Simpan session pengguna.
	session.SetSession(c, "user", strconv.Itoa(user.Id))

	return c.Redirect(http.StatusFound, "/index")
}

// LoginUser adalah fungsi untuk melakukan proses login melalui JSON data.
func LoginUser(c echo.Context) error {
	loginData := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	c.Bind(&loginData)

	// Cari pengguna berdasarkan email yang diberikan.
	user, err := repositories.GetAkunByEmail(loginData.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Email atau kata sandi salah")
	}

	// Bandingkan kata sandi yang dimasukkan dengan kata sandi yang ada di database.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Email atau kata sandi salah")
	}

	return c.JSON(http.StatusOK, "Login berhasil")
}

// Index adalah fungsi untuk menampilkan halaman index dari aplikasi.
func Index(c echo.Context) error {
	return c.File("config/views/index.html")
}