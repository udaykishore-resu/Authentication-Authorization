package middleware

import (
	"net/http"
	"session-cookie/handlers"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		sm := handlers.NewSessionManager()

		session, _ := sm.Store.Get(r, "session-name")

		if session.Values["user_id"] == nil {
			http.Error(rw, "Unathorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
