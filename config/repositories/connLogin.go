package repositories

import (
	"fmt"

	"part-1/config/models"
	"part-1/config/storage"
)

func GetAkunByEmail(email string) (*models.User, error) {
	db := storage.GetDB()
	user := &models.User{}
	sqlStatement := `SELECT id, nama, password, email FROM akun WHERE email = $1`

	err := db.QueryRow(sqlStatement, email).Scan(&user.Id, &user.Nama, &user.Password, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return user, nil
}