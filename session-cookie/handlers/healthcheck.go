package handlers

import (
	"encoding/json"
	"net/http"
	"session-cookie/database"
)

func HealthCheckHandler(db *database.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			http.Error(rw, "DB is not Pinging", http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]string{
			"message": "DB health is good. DB Pinned Successfully!!!",
		})
	}
}
