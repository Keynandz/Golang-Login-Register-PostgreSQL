package main

import (
	"part-1/config/handlers"
	"part-1/config/middleware"
	"part-1/config/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	// Membuat instance baru dari Echo
	e := echo.New()

	// Inisialisasi koneksi database
	storage.InitDB()

	// Mengatur route untuk beranda
	e.GET("/", handlers.Home)

	// Mengatur route untuk halaman indeks dengan middleware RequireLogin dan NoCache
	e.GET("/index", handlers.Index, middleware.RequireLogin, middleware.NoCache)

	// Mengatur route untuk logout
	e.GET("/exit", handlers.HandleLogout)

	// Mengatur route untuk register user
	e.POST("/register", handlers.HandleUser)

	// Mengatur route untuk login user
	e.POST("/login", handlers.HandleLogin)

	// Mengatur route untuk menampilkan halaman register (form)
	e.File("/form", "config/views/register.html")

	// Mengatur route untuk menampilkan halaman login
	e.File("/login", "config/views/login.html")

	// Mengatur route untuk menampilkan halaman logout
	e.File("/logout", "config/views/logout.html")

	// Mengatur route untuk menyajikan file statis dari folder views
	e.Static("/", "config/views")

	// Memulai server Echo pada port 5000
	e.Logger.Fatal(e.Start(":5000"))
}
