package main

import (
	"log"
	"my-go-web-server/src/connections"
	"my-go-web-server/src/models"
	"my-go-web-server/src/routes"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db := connections.ConnectToDB()
	db.AutoMigrate(&models.DBUser{})
	router := routes.SetupRoutes()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
