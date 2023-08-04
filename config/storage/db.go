package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

// GetDB mengembalikan instance database yang sudah terkoneksi.
func GetDB() *sql.DB {
	return db
}

// InitDB menginisialisasi koneksi ke database berdasarkan nilai dari file .env.
func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal mengambil data dari .env!")
	}

	// Ambil konfigurasi koneksi database dari file .env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Membuka koneksi ke database menggunakan driver postgres
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort))

	if err != nil {
		panic(err.Error())
	}

	// Melakukan ping ke database untuk memastikan koneksi berhasil
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Sukses! Terhubung ke database")
}