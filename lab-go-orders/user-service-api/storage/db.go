package storage

import (
	"log"
	"user-service-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(dsn string) {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return db
}
