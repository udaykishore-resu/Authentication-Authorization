package session

import (
	"session-cookie/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	User      *models.User
	CreatedAt time.Time
}

type SessionManager struct {
	sessions map[string]*Session
	mutex    sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*Session),
	}
}

func (sm *SessionManager) CreateSession(user *models.User) string {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	sessionID := uuid.New().String()
	sm.sessions[sessionID] = &Session{
		User:      user,
		CreatedAt: time.Now(),
	}

	return sessionID
}

func (sm *SessionManager) ValidateSession(sessionID string) bool {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	session, exists := sm.sessions[sessionID]
	if !exists {
		return false
	}

	// Optional: Add session expiration check
	if time.Since(session.CreatedAt) > 24*time.Hour {
		delete(sm.sessions, sessionID)
		return false
	}

	return true
}

func (sm *SessionManager) DeleteSession(sessionID string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	delete(sm.sessions, sessionID)
}
