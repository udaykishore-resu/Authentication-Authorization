package server

import (
	"json-web-token/auth"
	"json-web-token/config"
	"json-web-token/handlers"
	"json-web-token/repository"
	"net/http"
)

type Server struct {
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{config: cfg}
}

func (s *Server) Start() error {
	userRepo := repository.NewUserRepository()
	jwtService := auth.NewJWTService(s.config.JWTKey)
	authHandler := handlers.NewAuthHandler(userRepo, jwtService)

	// Public endpoints
	http.HandleFunc("/healthcheck", handlers.HealthCheck)
	http.HandleFunc("/login", authHandler.Login)

	// Protected endpoints
	http.HandleFunc("/logout", auth.AuthMiddleware(jwtService, userRepo)(authHandler.Logout))

	return http.ListenAndServe(s.config.Port, nil)
}
