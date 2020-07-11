package repository

import (
	"main/internal/models"
	"main/internal/tools/errors"
)


func (db *DB) Register(user *models.User) (*models.User, error) {
	err := db.DBConnPool.QueryRow(`INSERT INTO users (email, phone, password)
    VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id`, user.Email, user.Phone, user.Password).
	Scan(&user.ID)
	if err != nil {
		return nil, errors.UserAlreadyExists
	}
	return user, nil
}


func (db *DB) DeleteAccount(userID int) error {

	result, err := db.DBConnPool.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
	   return err
	}
	if result.RowsAffected() == 0 {
		return errors.UserNotFound
	}
	return nil
}
