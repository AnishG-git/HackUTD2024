package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AnishG-git/HackUTD2024/backend/internal/logic"
	"github.com/AnishG-git/HackUTD2024/backend/internal/models"
	"github.com/AnishG-git/HackUTD2024/backend/internal/util"
)

func (s *Server) registerHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a RegisterRequest object
	var reqBodyObj = *util.ExtractRequestBody[*models.RegisterRequest](r)

	// Perform register logic
	db := s.DB
	err := logic.Register(db, reqBodyObj)
	log.Println("testing out register logic")
	if err != nil {
		s.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Registration was successful
	const successfulRegistration = "user successfully created and inserted into DB"

	// Create a response object
	resBodyObj := models.RegisterResponse{
		Result: successfulRegistration,
	}

	// Encode the response object into JSON and write it to the response
	err = json.NewEncoder(w).Encode(resBodyObj)
	if err != nil {
		s.writeErrorResponse(w, http.StatusInternalServerError, encodingResponseError)
		return
	}
}
