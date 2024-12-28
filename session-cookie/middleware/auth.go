package middleware

import (
	"net/http"
	"session-cookie/session"
)

func AuthMiddleware(next http.HandlerFunc, sessionManager *session.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !sessionManager.ValidateSession(cookie.Value) {
			http.Error(w, "Invalid session", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
