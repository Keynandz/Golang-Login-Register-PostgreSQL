package main

import (
	"part-1/config/handlers"
	"part-1/config/middleware"
	"part-1/config/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	storage.InitDB()
	e.GET("/", handlers.Home)
	e.GET("/index", handlers.Index, middleware.RequireLogin, middleware.NoCache)
	e.GET("/exit", handlers.HandleLogout)

	e.POST("/register", handlers.HandleUser)
	e.POST("/login", handlers.HandleLogin)

	e.File("/form", "config/views/register.html")
	e.File("/login", "config/views/login.html")
	e.File("/logout", "config/views/logout.html")

	e.Static("/", "config/views")
	e.Logger.Fatal(e.Start(":5000"))
}
