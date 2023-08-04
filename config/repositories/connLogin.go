package repositories

import (
	"fmt"

	"part-1/config/models"
	"part-1/config/storage"
)

// GetAkunByEmail mengambil data akun dari database berdasarkan alamat email yang diberikan dan mengembalikan pointer ke data akun.
func GetAkunByEmail(email string) (*models.User, error) {
	// Dapatkan instance database yang sudah terkoneksi
	db := storage.GetDB()
	// Buat objek pointer untuk menyimpan data akun
	user := &models.User{}
	// SQL statement untuk mengambil data akun berdasarkan email
	sqlStatement := `SELECT id, nama, password, email FROM akun WHERE email = $1`

	// Eksekusi SQL statement dengan menggunakan QueryRow untuk mengambil data akun berdasarkan email
	err := db.QueryRow(sqlStatement, email).Scan(&user.Id, &user.Nama, &user.Password, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data akun: %w", err)
	}

	// Mengembalikan pointer ke data akun yang telah diambil dari database
	return user, nil
}