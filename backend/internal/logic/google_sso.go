package logic

import (
	"database/sql"

	"github.com/AnishG-git/HackUTD2024/backend/internal/util"
	"github.com/markbates/goth"
)

func GoogleSSO(db *sql.DB, user goth.User) (*string, error) {
	const (
		userDoesNotExist   = "user does not exist"
		insertUser         = "INSERT INTO users (email) VALUES ($1) RETURNING id"
		getUserCredentials = "SELECT id FROM users WHERE email = $1"
	)

	var userID string
	row := db.QueryRow(getUserCredentials, user.Email)
	err := row.Scan(&userID)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		// store user in db
		err = db.QueryRow(insertUser, user.Email).Scan(&userID)
		if err != nil {
			return nil, err
		}
	}

	jwt := util.GenerateToken(userID)

	// successful login
	return &jwt, nil
}
