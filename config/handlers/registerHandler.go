package handlers

import (
	"part-1/config/models"
	"part-1/config/repositories"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

func HandleUser(c echo.Context) error {
	contentType := c.Request().Header.Get("Content-Type")
	if contentType == "application/json" {
		return CreateAkun(c)
	}
	return CreateAkunWeb(c)
}

// FUNC REGISTER USER BY WEB
func CreateAkunWeb(c echo.Context) error {
	nama := c.FormValue("nama")
	password := c.FormValue("password")
	email := c.FormValue("email")

	hashedPassword, err := encryptPasswordWeb(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := models.User{
		Nama: nama,
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

// FUNC REGISTER USER
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

// FUNC ENCRYPT PASS
func encryptPasswordWeb(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func encryptPassword(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
