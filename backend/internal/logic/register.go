package logic

import (
	"database/sql"
	"fmt"

	"github.com/AnishG-git/HackUTD2024/backend/internal/models"
)

func Register(db *sql.DB, reqBody models.RegisterRequest) error {
	const (
		createNewUserQuery      = `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
		failedToStoreUserInDB   = "failed to store user in DB"
		userAlreadyExists       = "user already exists"
		userCreatedSuccessfully = "user successfully created and inserted into DB"
	)

	username := *reqBody.Username
	email := *reqBody.Email
	password := *reqBody.Password

	row := db.QueryRow(getExistingUserQuery, reqBody.Email)
	
	var existingUserID string
	err := row.Scan(&existingUserID)

	// any error that is not sql.ErrNoRows is unexpected
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// no error means user already exists
	if err == nil {
		err = fmt.Errorf(userAlreadyExists)
		return err
	}

	// user does not exist so create a new user
	_, err = db.Exec(createNewUserQuery, username, email, password)
	if err != nil {
		return err
	}

	// successfully created user
	return nil
}
