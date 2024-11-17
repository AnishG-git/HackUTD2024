package logic

import (
	"database/sql"

	"github.com/AnishG-git/HackUTD2024/backend/internal/util"
	"github.com/google/uuid"
	"github.com/markbates/goth"
)

func GoogleSSO(db *sql.DB, user goth.User) (*string, error) {
	const (
		userDoesNotExist   = "user does not exist"
		insertUser         = "INSERT INTO users (email, sdk_key) VALUES ($1, $2) RETURNING sdk_key"
		getUserCredentials = "SELECT sdk_key FROM users WHERE email = $1"
	)

	var sdk_key string
	row := db.QueryRow(getUserCredentials, user.Email)
	err := row.Scan(&sdk_key)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		sdk_key = uuid.New().String()
		// store user in db
		err = db.QueryRow(insertUser, user.Email, sdk_key).Scan(&sdk_key)
		if err != nil {
			return nil, err
		}
	}

	jwt := util.GenerateToken(sdk_key)

	// successful login
	return &jwt, nil
}
