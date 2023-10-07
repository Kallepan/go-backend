package repository

import (
	"database/sql"

	"github.com/kallepan/go-backend/app/domain/dao"
)

type UserRepository interface {
	RegisterUser(user *dao.User) (string, error)
	CheckIfUserExists(username string) bool
	GetUserByUsername(username string) (*dao.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func (u UserRepositoryImpl) GetUserByUsername(username string) (*dao.User, error) {
	/* Returns the user object with the user_id */
	var user dao.User

	query := `
		SELECT user_id, email, firstname, lastname, username, password
		FROM users
		WHERE username = $1`
	err := u.db.QueryRow(query, username).Scan(
		&user.UserId,
		&user.Email,
		&user.Firstname,
		&user.Lastname,
		&user.Username,
		&user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserRepositoryImpl) RegisterUser(user *dao.User) (string, error) {
	/* Creates the user and returns the user object with the user_id */

	// Insert user into database
	query := `
		INSERT INTO users (email, firstname, lastname, username, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id`
	var userID string
	err := u.db.QueryRow(
		query,
		&user.Email,
		&user.Firstname,
		&user.Lastname,
		&user.Username,
		&user.Password).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func UserRepositoryInit(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) CheckIfUserExists(username string) bool {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE username = $1
		);`

	err := u.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}
