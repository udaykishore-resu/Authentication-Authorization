package database

import (
	"database/sql"
	"errors"
	"session-cookie/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Authenticate(username, password string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username FROM user WHERE username = $1 AND password = $2`
	err := r.db.QueryRow(query, username, password).Scan(&user.ID, &user.Username)

	if err == sql.ErrNoRows {
		return nil, errors.New("invalid credentials")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
