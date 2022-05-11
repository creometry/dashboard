package main

import (
	"log"

	"github.com/Creometry/dashboard/auth"
	"github.com/Creometry/dashboard/routes"
	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func main() {
	// Create the client set
	auth.CreateKubernetesClient()

	// Get config from .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	routes.CreateRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
