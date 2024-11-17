package main

import (
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

func newServer() *mux.Router {
	server := mux.NewRouter()
	api := server.PathPrefix("/api").Subrouter()
	api.HandleFunc("/post-request-1", postOneHandler).Methods("POST")
	return server
}

func postOneHandler(w http.ResponseWriter, r *http.Request) {
	wheel := rand.Intn(2)
	if wheel == 0 {

	} else {

	}
}
