package database

import (
	"go-api-car/models"
	"log"
	"os"

	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Connection() {
	conn := "host=localhost user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_DATABASE") + " port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conn))

	if err != nil {
		log.Panic("Database connection error")
	}

	DB.AutoMigrate(&models.Car{})	
}
