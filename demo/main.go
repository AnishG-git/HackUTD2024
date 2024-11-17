package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/AnishG-git/HackUTD2024/backend/pineapplesdk"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// allowing cross origin requests for local testing
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	sdk_key := "47e05264-1e37-416b-b4d2-00e212ab1582"

	server := newServer()
	server.Use(pineapplesdk.CreateMW(sdk_key))
	go func() {
		err := http.ListenAndServe("localhost:3001", cors(server))
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
		log.Println("Server started on port 7000")
	}()

	time.Sleep(1 * time.Second)

	// barrage the server with 10 requests
	// barrageServer(10)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:3001/api/get-handler-1", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	barrageServer(10)
	time.Sleep(2 * time.Second)
}

func barrageServer(numRequests int) {
	wg := sync.WaitGroup{}
	options := []string{"http://localhost:3001/api/post-handler-1", "http://localhost:3001/api/post-handler-2", "http://localhost:3001/api/post-handler-3", "http://localhost:3001/api/get-handler-1", "http://localhost:3001/api/get-handler-2"}
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			randomNum := rand.Intn(len(options))
			randomURL := options[randomNum]
			client := &http.Client{}
			fmt.Println(randomURL)
			var req *http.Request
			if strings.Contains(randomURL, "post-handler") {
				tmp, err := http.NewRequest("POST", randomURL, strings.NewReader(`{"message":"the random number is `+string(randomNum)+`"}`))
				if err != nil {
					log.Fatalf("Error creating request: %v", err)
				}
				req = tmp
			} else {
				tmp, err := http.NewRequest("GET", randomURL, nil)
				if err != nil {
					log.Fatalf("Error creating request: %v", err)
				}
				req = tmp
			}
			req.Header.Set("Content-Type", "application/json")
			_, err := client.Do(req)
			if err != nil {
				log.Fatalf("Error sending request: %v", err)
			}
			// time.Sleep(5 * time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
}

// Response structure for success or error
type response struct {
	Message string `json:"message"`
}

// Create a new server and define the routes
func newServer() *mux.Router {
	server := mux.NewRouter()
	api := server.PathPrefix("/api").Subrouter()

	// POST handlers
	api.HandleFunc("/post-handler-1", postHandler1).Methods("POST")
	api.HandleFunc("/post-handler-2", postHandler2).Methods("POST")
	api.HandleFunc("/post-handler-3", postHandler3).Methods("POST")

	// GET handlers
	api.HandleFunc("/get-handler-1", getHandler1).Methods("GET")
	api.HandleFunc("/get-handler-2", getHandler2).Methods("GET")

	return server
}

// POST handler 1
func postHandler1(w http.ResponseWriter, r *http.Request) {
	processRequest(w, r, "Handler 1")
}

// POST handler 2
func postHandler2(w http.ResponseWriter, r *http.Request) {
	processRequest(w, r, "Handler 2")
}

// POST handler 3
func postHandler3(w http.ResponseWriter, r *http.Request) {
	processRequest(w, r, "Handler 3")
}

// GET handler 1
func getHandler1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Message: "GET Handler 1 - Data retrieved successfully",
	})
}

// GET handler 2
func getHandler2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Message: "GET Handler 2 - Resource found",
	})
}

// Process request logic for POST handlers
func processRequest(w http.ResponseWriter, r *http.Request, handlerName string) {
	// Seed random generator
	wheel := rand.Intn(2) // Generate random number: 0 or 1

	// Decode the JSON body
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Simulate success or error based on random number
	if wheel == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{
			Message: handlerName + " - Success",
		})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{
			Message: handlerName + " - Error",
		})
	}
}
