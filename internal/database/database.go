package database

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"

	"github.com/vlbarou/sampleproject/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable SQL loggin
	})
	if err != nil {
		log.Fatalf("Failed to connect to MariaDB: %v", err)
	}

	return db
}

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
	sqlDB.Close()
}
