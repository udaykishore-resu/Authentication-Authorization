package handlers

import (
	"encoding/json"
	"net/http"
	"session-cookie/database"
	"session-cookie/logger"
	"session-cookie/session"
	"time"
)

type Handler struct {
	userRepo       *database.UserRepository
	sessionManager *session.SessionManager
	logger         *logger.Logger
}

func NewHandler(
	userRepo *database.UserRepository,
	sessionManager *session.SessionManager,
	logger *logger.Logger,
) *Handler {
	return &Handler{
		userRepo:       userRepo,
		sessionManager: sessionManager,
		logger:         logger,
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.Authenticate(req.Username, req.Password)
	if err != nil {
		h.logger.Error("Authentication failed")

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Create session
	sessionID := h.sessionManager.CreateSession(user)

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionID,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "No session found", http.StatusBadRequest)
		return
	}

	h.sessionManager.DeleteSession(cookie.Value)

	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "healthy",
		"message": "Database connection is active",
	})
}
