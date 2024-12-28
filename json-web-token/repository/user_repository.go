package repository

import (
	"sync"
)

type UserRepository struct {
	users             map[string]string
	blacklistedTokens struct {
		sync.RWMutex
		tokens map[string]bool
	}
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: map[string]string{
			"admin": "admin123", // In production, store hashed passwords
		},
		blacklistedTokens: struct {
			sync.RWMutex
			tokens map[string]bool
		}{tokens: make(map[string]bool)},
	}
}

func (r *UserRepository) ValidateCredentials(username, password string) bool {
	expectedPassword, ok := r.users[username]
	return ok && expectedPassword == password
}

func (r *UserRepository) BlacklistToken(token string) {
	r.blacklistedTokens.Lock()
	defer r.blacklistedTokens.Unlock()
	r.blacklistedTokens.tokens[token] = true
}

func (r *UserRepository) IsTokenBlacklisted(token string) bool {
	r.blacklistedTokens.RLock()
	defer r.blacklistedTokens.RUnlock()
	return r.blacklistedTokens.tokens[token]
}
