package main

import (
	"log"

	"github.com/Creometry/dashboard/auth"
	"github.com/Creometry/dashboard/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/joho/godotenv"
)

func main() {
	// Create the client set
	auth.CreateKubernetesClient()
	auth.CreateExtensionsClient()

	// Get config from .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// add CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.CreateRoutes(app)

	log.Fatal(app.Listen(":3001"))
}
