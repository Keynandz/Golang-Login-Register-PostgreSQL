package handlers

import (
	"part-1/config/repositories"
	"part-1/config/session"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(c echo.Context) error {
	contentType := c.Request().Header.Get("Content-Type")
	if contentType == "application/json" {
		return LoginUser(c)
	}
	return LoginUserWeb(c)
}

func LoginUserWeb(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := repositories.GetAkunByEmail(email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}

	session.SetSession(c, "user", strconv.Itoa(user.Id))

	return c.Redirect(http.StatusFound, "/index")
}

func LoginUser(c echo.Context) error {
	loginData := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	c.Bind(&loginData)

	user, err := repositories.GetAkunByEmail(loginData.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}

	return c.JSON(http.StatusOK, "Login successful")
}

func Index(c echo.Context) error {
	return c.File("config/views/index.html")
}
