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
	// this is just for local development
	// in production, create a restCLient without the use os a kubeconfig
	// file. This is because the kubeconfig file is not available in the
	// container. Will use a service account instead.
	// (https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration)
	auth.CreateKubernetesClient()
	auth.CreateExtensionsClient()

	// Get config from .env file
	// this is just for local development
	// in production, the app gets the enviroment variables from the configmap
	// in the cluster.
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// add CORS
	// this is just for local development
	// the allowed origin should be the service name of the frontend
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.CreateRoutes(app)

	log.Fatal(app.Listen(":3001"))
}
