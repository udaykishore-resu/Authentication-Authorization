package auth

import (
	"json-web-token/repository"
	"net/http"
	"strings"
)

func AuthMiddleware(jwtService *JWTService, userRepo *repository.UserRepository) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 {
				http.Error(w, "Invalid token format", http.StatusUnauthorized)
				return
			}

			tokenStr := bearerToken[1]

			if userRepo.IsTokenBlacklisted(tokenStr) {
				http.Error(w, "Token is blacklisted", http.StatusUnauthorized)
				return
			}

			token, err := jwtService.ValidateToken(tokenStr)
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		}
	}
}
