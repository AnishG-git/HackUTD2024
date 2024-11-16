package logic

import (
	"database/sql"
	"io"
	"log"
	"net/http"
)

func Predict(db *sql.DB, logger *log.Logger, reqBody io.Reader) ([]byte, error) {
	client := http.Client{}
	resp, err := client.Post("http://model:8000/predict", "application/json", reqBody)
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return respBytes, nil
}
