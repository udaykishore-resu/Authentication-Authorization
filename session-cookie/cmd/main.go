package main

import (
	"log"
	"net/http"
	"session-cookie/database"
	"session-cookie/handlers"
	"session-cookie/middleware"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}
	defer db.Close()

	sessionManager := handlers.NewSessionManager()

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handlers.LoginHandler(db, sessionManager))
	mux.HandleFunc("/logout", middleware.AuthMiddleware(handlers.LogoutHandler(sessionManager)))
	mux.HandleFunc("/healthcheck", middleware.AuthMiddleware(handlers.HealthCheckHandler(db)))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
