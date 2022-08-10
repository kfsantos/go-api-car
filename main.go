package main

import (
	"go-api-car/database"
	"go-api-car/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.Connection()
	r := routes.HandlerRequests()
	if err := r.Run("127.0.0.1:5000"); err != nil {
		log.Fatal(err)
	}
}
