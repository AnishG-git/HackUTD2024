package main

import (
	"log"
	"net/http"

	"github.com/AnishG-git/HackUTD2024/backend/internal/server"
	"github.com/AnishG-git/HackUTD2024/backend/storage"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

const (
	port = "8080"
)

func main() {
	mainLog := log.Default()
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	db, err := storage.LoadDB("hackutd2024", mainLog)
	if err != nil {
		mainLog.Fatalf("Failed to load database: %v", err)
	}

	srv := server.NewServer(port, db, mainLog)

	mainLog.Printf("Starting server on port %s\n", srv.Address)

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	err = http.ListenAndServe(":"+srv.Address, cors(srv.Router))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
