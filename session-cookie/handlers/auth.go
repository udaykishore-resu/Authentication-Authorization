package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"session-cookie/database"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

type SessionManager struct {
	Store *sessions.CookieStore
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		Store: sessions.NewCookieStore([]byte("secret-key")),
	}
}

func LoginHandler(db *database.DB, sm *SessionManager) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var creds struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(rw, "Invalid request body", http.StatusBadRequest)
			return
		}

		user, err := db.GetUser(creds.Username)
		if err != nil {
			if err.Error() == fmt.Sprintf("no user found with username: %s", creds.Username) {
				http.Error(rw, "Invalid credentials", http.StatusUnauthorized)
			} else {
				http.Error(rw, "Internal server error", http.StatusInternalServerError)
				log.Printf("Error getting user: %v", err)
			}
			return
		}

		if user.Password != creds.Password {
			http.Error(rw, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		session, _ := sm.Store.Get(r, "session-name")
		session.Values["user_id"] = user.ID
		session.Values["session_id"] = uuid.New().String()
		session.Options.MaxAge = 3600
		session.Save(r, rw)

		rw.WriteHeader(http.StatusOK)

		json.NewEncoder(rw).Encode(map[string]string{
			"message": "Logged in Successfully!!!",
		})
	}
}

func LogoutHandler(sm *SessionManager) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		session, _ := sm.Store.Get(r, "session-name")
		session.Options.MaxAge = -1
		session.Save(r, rw)
		rw.WriteHeader(http.StatusOK)

		json.NewEncoder(rw).Encode(map[string]string{
			"message": "Logged out Successfully ",
		})
	}
}
