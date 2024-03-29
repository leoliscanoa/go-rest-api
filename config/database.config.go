package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"rest-api/model"
)

func InitDatabase() *gorm.DB {
	jdbcUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(jdbcUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Println("Database connection successfully")
	errMigrate := db.AutoMigrate(&model.User{})
	if errMigrate != nil {
		panic("Failed to migrate models to config")
	}
	return db
}
