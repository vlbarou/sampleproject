package main

import (
	"log"

	"github.com/vlbarou/sampleproject/config"
	"github.com/vlbarou/sampleproject/internal/database"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)
	defer database.Close(db)

	if err := database.Migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully")
}
