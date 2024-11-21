package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

type User struct {
	ID       int
	Username string
	Password string
}

func ConnectDB() (*DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) GetUser(username string) (User, error) {
	var user User
	err := db.QueryRow("SELECT user_id, username, password FROM user WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// No user found with the given username
			return User{}, fmt.Errorf("no user found with username: %s", username)
		}
		// Some other error occurred
		return User{}, fmt.Errorf("error querying user: %v", err)
	}
	return user, nil
}
