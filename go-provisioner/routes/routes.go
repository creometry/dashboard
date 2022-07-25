package routes

import (
	pr "github.com/Creometry/dashboard/go-provisioner/controllers"
	gh "github.com/Creometry/dashboard/go-provisioner/controllers/github"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(app *fiber.App) {

	v1 := app.Group("/api/v1")
	v1.Get("/github/exchange/:code", gh.GetAccessToken)
	v1.Post("/provisionProject", pr.ProvisionProject)
	v1.Post("/login", pr.Login)
	v1.Post("/register", pr.Register)
	v1.Get("/team/:projectId", pr.ListTeamMembers)
	v1.Post("/team/:projectId/:userId", pr.AddTeamMember)
	v1.Post("/kubeconfig", pr.GenerateKubeConfig)
}
