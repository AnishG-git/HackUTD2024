package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/AnishG-git/HackUTD2024/backend/internal/util"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")

		// check for empty authorization token
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Missing authorization header"}`))
			return
		}

		// check for valid authorization token format
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Invalid authorization header"}`))
			return
		}

		tokenString = tokenParts[1]

		// verify jwt
		claims, err := util.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Unauthorized token"}`))
			return
		}

		// set user id on context for downstream handlers
		ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
		next(w, r.WithContext(ctx))
	}
}
