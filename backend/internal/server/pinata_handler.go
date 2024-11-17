package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AnishG-git/HackUTD2024/backend/internal/models"
	"github.com/AnishG-git/HackUTD2024/backend/internal/util"
)

func (s *Server) handlePinata(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a RegisterRequest object
	reqBodyObj := *util.ExtractRequestBody[*models.Pinata](r)
	// Construct the response object
	response := map[string]string{
		"gateway": fmt.Sprintf("https://moccasin-improved-smelt-48.mypinata.cloud/files/%s", *(reqBodyObj.Cid)),
	}

	// Set status code and encode the response as JSON
	w.WriteHeader(http.StatusOK) // Set status code
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
