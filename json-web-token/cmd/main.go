package main

import (
	"json-web-token/config"
	"json-web-token/server"
	"log"
)

func main() {
	cfg := config.NewConfig()
	srv := server.NewServer(cfg)

	log.Printf("Server starting on port %s...", cfg.Port)
	log.Fatal(srv.Start())
}
