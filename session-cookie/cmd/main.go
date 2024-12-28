package main

import (
	"log"
	"net/http"
	"session-cookie/config"
	"session-cookie/database"
	"session-cookie/handlers"
	"session-cookie/logger"
	"session-cookie/middleware"
	"session-cookie/session"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize configuration
	cfg := config.NewConfig()

	// Setup logger
	appLogger := logger.NewLogger()

	// Initialize database connection
	db, err := database.NewConnection(cfg)
	if err != nil {
		appLogger.Fatal("Database connection failed", "error", err)
	}
	defer db.Close()

	// Initialize session manager
	sessionManager := session.NewSessionManager()

	// Setup router
	router := mux.NewRouter()

	// Initialize auth handler
	authHandler := handlers.NewHandler(
		database.NewUserRepository(db),
		sessionManager,
		appLogger,
	)

	// Routes
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	router.HandleFunc("/healthcheck",
		middleware.AuthMiddleware(authHandler.HealthCheck, sessionManager),
	).Methods("GET")

	// Start server
	serverAddr := cfg.ServerAddress()
	appLogger.Info("Server starting", "address", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
