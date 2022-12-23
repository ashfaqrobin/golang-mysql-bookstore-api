package main

import (
	"log"

	"github.com/ashfaqrobin/golang-mysql-bookstore-api/pkg/config"
	"github.com/ashfaqrobin/golang-mysql-bookstore-api/pkg/routes"
	"github.com/ashfaqrobin/golang-mysql-bookstore-api/pkg/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variables
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Failed to load environment")
	}

	// Connect to database
	config.Connect()

	// Setup Repository
	repo := storage.CreateRepository()
	storage.SetRepository(repo)

	// Creating fiber app
	app := fiber.New()

	// Regester routes
	routes.RegisterBookStoreRoutes(app)

	// Listen to port
	app.Listen(":8000")
}
