package repositories

import (
	"fmt"

	"part-1/config/models"
	"part-1/config/storage"
)

// CreateAkun membuat akun baru dalam database berdasarkan data yang diberikan dan mengembalikan akun yang baru dibuat.
func CreateAkun(user models.User) (models.User, error) {
	// Dapatkan instance database yang sudah terkoneksi
	db := storage.GetDB()
	// SQL statement untuk memasukkan data akun ke dalam database dan mengembalikan id yang baru dibuat
	sqlStatement := `INSERT INTO akun (nama, email, password, verif) VALUES ($1, $2, $3, $4) RETURNING id`

	// Eksekusi SQL statement dengan menggunakan QueryRow untuk mendapatkan id yang baru ditambahkan
	err := db.QueryRow(sqlStatement, user.Nama, user.Email, user.Password, user.Verif).Scan(&user.Id)
	if err != nil {
		return user, fmt.Errorf("membuat akun baru error: %w", err)
	}

	// Mengembalikan akun yang baru dibuat
	return user, nil
}