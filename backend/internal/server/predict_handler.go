package server

import (
	"encoding/json"
	"net/http"

	"github.com/AnishG-git/HackUTD2024/backend/internal/logic"
)

func (s *Server) predictHandler(w http.ResponseWriter, r *http.Request) {
	db := s.DB
	logger := s.Logger
	result, err := logic.Predict(db, logger, r.Body)
	if err != nil {
		s.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// python model api result (could be success or failure)
	type modelResponse struct {
		Error string `json:"error"`
	}

	// checking if an error was received in response from python model api
	var modelRes modelResponse
	if err := json.Unmarshal(result, &modelRes); err != nil {
		s.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// attaching proper status code to response
	if modelRes.Error != "" {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(result)
}
