package main

import (
	"log"

	"github.com/Creometry/dashboard/go-provisioner/auth"
	"github.com/Creometry/dashboard/go-provisioner/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	auth.CreateInClusterClient()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.CreateRoutes(app)

	log.Fatal(app.Listen(":3001"))
}
