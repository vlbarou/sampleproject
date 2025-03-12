package main

import (
	"github.com/vlbarou/sampleproject/config"
	"github.com/vlbarou/sampleproject/internal/database"
	"log"

	"github.com/vlbarou/sampleproject/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)
	defer database.Close(db)

	srv := server.NewServer(cfg, db)
	log.Fatal(srv.Start())
}
