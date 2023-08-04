package handlers

import (
	"part-1/config/models"
	"part-1/config/repositories"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

// HandleUser adalah fungsi untuk menangani request dari endpoint user.
func HandleUser(c echo.Context) error {
	contentType := c.Request().Header.Get("Content-Type")
	if contentType == "application/json" {
		return CreateAkun(c)
	}
	return CreateAkunWeb(c)
}

// CreateAkunWeb adalah fungsi untuk mendaftarkan pengguna melalui web dengan form values.
func CreateAkunWeb(c echo.Context) error {
	nama := c.FormValue("nama")
	password := c.FormValue("password")
	email := c.FormValue("email")

	hashedPassword, err := encryptPasswordWeb(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := models.User{
		Nama:     nama,
		Password: hashedPassword,
		Email:    email,
		Verif:    false,
	}

	_, err = repositories.CreateAkun(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusFound, "/login")
}

// CreateAkun adalah fungsi untuk mendaftarkan pengguna dengan data JSON.
func CreateAkun(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := encryptPassword(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	newUser, err := repositories.CreateAkun(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newUser)
}

// encryptPasswordWeb adalah fungsi untuk mengenkripsi password dari registrasi melalui web.
func encryptPasswordWeb(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// encryptPassword adalah fungsi untuk mengenkripsi password dari registrasi melalui JSON.
func encryptPassword(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}