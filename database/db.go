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
	db, err := gorm.Open(
		os.Getenv("DB_CONNECTION"),
		"host="+os.Getenv("DB_HOST")+
			" port="+os.Getenv("DB_PORT")+
			" user="+os.Getenv("DB_USERNAME")+
			" dbname="+os.Getenv("DB_DATABASE")+
			" sslmode=disable password="+os.Getenv("DB_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}
	//Para visualizar consultas SQL
	db.LogMode(true)
	//cria a tabela chamada books
	db.AutoMigrate([]models.Car{})
	DB = db
}
