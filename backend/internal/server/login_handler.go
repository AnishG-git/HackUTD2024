package server

import (
	"encoding/json"
	"net/http"

	"github.com/AnishG-git/HackUTD2024/backend/internal/logic"
	"github.com/AnishG-git/HackUTD2024/backend/internal/models"
	"github.com/AnishG-git/HackUTD2024/backend/internal/util"
)

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	
	var reqBodyObj = *util.ExtractRequestBody[*models.LoginRequest](r)

	db := s.DB
	jwt, err := logic.Login(db, reqBodyObj)
	if err != nil {
		s.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	respBodyObj := models.LoginResponse{
		JWT: *jwt,
	}

	err = json.NewEncoder(w).Encode(respBodyObj)
	if err != nil {
		s.writeErrorResponse(w, http.StatusInternalServerError, encodingResponseError)
		return
	}
}
