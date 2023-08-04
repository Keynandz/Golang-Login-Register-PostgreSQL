package repositories

import (
	"fmt"

	"part-1/config/models"
	"part-1/config/storage"
)

func CreateAkun(user models.User) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO akun (nama, email, password) VALUES ($1, $2, $3) RETURNING id`

	err := db.QueryRow(sqlStatement, user.Nama, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return user, fmt.Errorf("membuat akun baru error: %w", err)
	}

	return user, nil
}