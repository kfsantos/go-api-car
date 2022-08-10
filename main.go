package main

import (
	"go-api-car/database"
	"go-api-car/routes"
	"log"
)

func main() {
	database.Connection()
	r := routes.HandlerRequests()
	if err := r.Run("127.0.0.1:5000"); err != nil {
		log.Fatal(err)
	}
}
