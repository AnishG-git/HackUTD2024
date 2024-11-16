package server

import (
	"net/http"

	"github.com/AnishG-git/HackUTD2024/backend/internal/logic"
	"github.com/markbates/goth/gothic"
)

func (s *Server) googleAuthHandler(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (s *Server) googleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	var jwt *string
	if user, err := gothic.CompleteUserAuth(w, r); err == nil {
		s.Logger.Printf("User %s signing in", user.Email)
		jwt, err = logic.GoogleSSO(s.DB, user)
		if err != nil {
			s.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		s.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		// http.Redirect(w, r, "http://localhost:3000", http.StatusTemporaryRedirect)
		return
	}

	// type response struct {
	// 	JWT string `json:"jwt"`
	// }

	// session, _ := gothic.Store.Get(r, gothic.SessionName)
	// session.Values["jwt"] = *jwt
	// session.Save(r, w)

	// log jwt
	s.Logger.Printf("JWT: %s", *jwt)
	// json.NewEncoder(w).Encode(response{JWT: *jwt})
	c := &http.Cookie{
		Name:     "jwt",
		Value:    *jwt,
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   86400 * 30, // 30 days
		HttpOnly: false,      // Make sure the cookie is not accessible via JavaScript
		Secure:   false,      // Set to true if using HTTPS
	}

	http.SetCookie(w, c)
	http.Redirect(w, r, "http://localhost:3000", http.StatusTemporaryRedirect)
}

func (s *Server) googleLogoutHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// // jwtCookie, err := r.Cookie("jwt")
	// // if err != nil {
	// // 	s.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
	// // 	return
	// // }
	s.Logger.Printf("Logging out\n")
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false, // Make sure the cookie is not accessible via JavaScript
		Secure:   false, // Set to true if using HTTPS
	}

	http.SetCookie(w, cookie)

	// err := gothic.Logout(w, r)
	// if err != nil {
	// 	s.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	http.Redirect(w, r, "http://localhost:3000", http.StatusTemporaryRedirect)
}
