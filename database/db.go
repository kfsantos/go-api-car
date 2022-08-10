package database

import (
	"go-api-car/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connection() {
	conn := "host=localhost user=postgres password=1234 dbname=teste port=5432 sslmod=disable"
	DB, err := gorm.Open(postgres.Open(conn))

	if err != nil {
		log.Panic("Database connection error")
	}
	DB.AutoMigrate(&models.Car{})

	// db, err := gorm.Open(
	// 	os.Getenv("DB_CONNECTION"),
	// 		"host="+os.Getenv("DB_HOST")+
	// 		" port="+os.Getenv("DB_PORT")+
	// 		" user="+os.Getenv("DB_USERNAME")+
	// 		" dbname="+os.Getenv("DB_DATABASE")+
	// 		" sslmode=disable password="+os.Getenv("DB_PASSWORD"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //Para visualizar consultas SQL
	// db.Logger.LogMode(true)
	// //cria a tabela chamada books
	// db.AutoMigrate([]models.Car{})
	// DB = db
}
