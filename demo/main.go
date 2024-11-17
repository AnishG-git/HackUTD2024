package main

import (
	"log"
	"net/http"

	"github.com/AnishG-git/HackUTD2024/backend/pineapplesdk"
	"github.com/gorilla/handlers"
)

func main() {
	// allowing cross origin requests for local testing
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	sdk_key := ""
	pineappleMW := pineapplesdk.CreateMW(sdk_key)
	server := newServer()
	server.Use(pineappleMW)

	go func() {
		err := http.ListenAndServe(":7000", cors(server))
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}
