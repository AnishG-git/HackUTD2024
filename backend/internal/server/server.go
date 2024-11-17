package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	mw "github.com/AnishG-git/HackUTD2024/backend/internal/middleware"
	"github.com/AnishG-git/HackUTD2024/backend/internal/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type Server struct {
	Router  *mux.Router
	Address string
	DB      *sql.DB
	Logger  *log.Logger
	Store   *sessions.CookieStore
}

func NewServer(addr string, db *sql.DB, logger *log.Logger) *Server {
	// initialize gothic for sso
	store := sessions.NewCookieStore([]byte("secret"))
	store.MaxAge(86400 * 30)
	gothic.Store = store
	clientKey := os.Getenv("CLIENT_KEY")
	clientSecret := os.Getenv("CLIENT_SECRET")
	goth.UseProviders(
		google.New(clientKey, clientSecret, "http://localhost:8080/api/auth/google/callback", "email", "profile"),
	)

	router := mux.NewRouter()
	// router.Use(pineapplesdk.CreateMW("47e05264-1e37-416b-b4d2-00e212ab1582"))

	s := &Server{
		Router:  router,
		Address: addr,
		DB:      db,
		Logger:  logger,
		Store:   store,
	}

	s.routes()
	return s
}

type middleware func(http.HandlerFunc) http.HandlerFunc

// wrapHandler applies multiple middlewares to the handler
func wrapHandler(h http.HandlerFunc, mws ...middleware) http.HandlerFunc {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

func attachDefaultMws[ReqModel any](h http.HandlerFunc) http.HandlerFunc {
	return wrapHandler(h, mw.JsonMiddleware[ReqModel])
}

func attachDefaultMwsWithAuth[ReqModel any](h http.HandlerFunc) http.HandlerFunc {
	return wrapHandler(h, mw.JsonMiddleware[ReqModel], mw.AuthMiddleware)
}

func (s *Server) routes() {

	// handlers that don't require authentication
	registerHandler := attachDefaultMws[models.RegisterRequest](s.registerHandler)
	loginHandler := attachDefaultMws[models.LoginRequest](s.loginHandler)
	pinataHandler := attachDefaultMws[models.Pinata](s.handlePinata)

	// handlers that require authentication
	predictHandler := attachDefaultMwsWithAuth[models.PredictRequest](s.predictHandler)

	// API routes
	api := s.Router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/register", registerHandler).Methods("POST")
	api.HandleFunc("/login", loginHandler).Methods("POST")
	api.HandleFunc("/predict", predictHandler).Methods("POST")
	api.HandleFunc("/fetch-data/{sdk-key}", s.fetchDataHandler).Methods("GET")
	api.HandleFunc("/pinata", pinataHandler).Methods("POST")

	// sso routes
	authR := api.PathPrefix("/auth/{provider}").Subrouter()
	authR.HandleFunc("/sso", s.googleAuthHandler).Methods("GET")
	authR.HandleFunc("/callback", s.googleCallbackHandler).Methods("GET")
	authR.HandleFunc("/logout", s.googleLogoutHandler).Methods("GET")
}

// writeErrorResponse writes an error response to the client in JSON and logs the error
func (s *Server) writeErrorResponse(w http.ResponseWriter, statusCode int, err string) {
	s.Logger.Printf("error: %s", err)
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
}
