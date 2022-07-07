package main

import (
	"log"

	"github.com/Creometry/resources-service/auth"
	"github.com/Creometry/resources-service/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/joho/godotenv"
)

func main() {
	auth.CreateKubernetesClient()
	auth.CreateExtensionsClient()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.CreateRoutes(app)

	log.Fatal(app.Listen(":3002"))
}
