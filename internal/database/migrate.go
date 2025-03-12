package database

import (
	"log"

	"golang/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Println("Running database migrations...")
	return db.AutoMigrate(&model.User{})
}
