package handlers

import (
	"fmt"
	"net/http"

	"part-1/config/session"

	"github.com/labstack/echo/v4"
)

func HandleLogout(c echo.Context) error {
	session.ClearSession(c)
	fmt.Println("clearing session")

	return c.Redirect(http.StatusFound, "/logout")
}
