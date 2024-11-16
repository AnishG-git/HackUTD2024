package logic

import (
	"database/sql"
	"fmt"

	"github.com/AnishG-git/HackUTD2024/backend/internal/models"
	"github.com/AnishG-git/HackUTD2024/backend/internal/util"
)

func Login(db *sql.DB, reqBody models.LoginRequest) (*string, error) {
	const (
		userDoesNotExist   = "user does not exist"
		getUserCredentials = "SELECT id, password FROM users WHERE email = $1"
		invalidCredentials = "invalid credentials"
	)

	// password sent in login request
	password := *reqBody.Password

	var existingUserID string
	var storedPassword string

	row := db.QueryRow(getUserCredentials, *reqBody.Email)
	err := row.Scan(&existingUserID, &storedPassword)

	// user does not exist
	if err != nil && err == sql.ErrNoRows {
		return nil, fmt.Errorf(userDoesNotExist)
	}

	// unexpected error
	if err != nil {
		return nil, err
	}

	// authenticate user
	if password != storedPassword {
		return nil, fmt.Errorf(invalidCredentials)
	}

	// generate jwt token
	jwt := util.GenerateToken(existingUserID)

	// successful login
	return &jwt, nil
}
