package main

import (
	"golang/config"
	"golang/internal/database"
	"log"

	"golang/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)
	defer database.Close(db)

	srv := server.NewServer(cfg, db)
	log.Fatal(srv.Start())
}
