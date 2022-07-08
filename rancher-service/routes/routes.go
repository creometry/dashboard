package routes

import (
	gh "github.com/Creometry/rancher-service/controllers/github"
	pr "github.com/Creometry/rancher-service/controllers/project"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(app *fiber.App) {

	v1 := app.Group("/api/v1")
	v1.Get("/github/exchange/:code", gh.GetAccessToken)
	v1.Post("/project", pr.CreateProject)
	v1.Get("/user/:username", pr.FindUserAndLoginOrCreate)
	v1.Get("/team/:projectId", pr.ListTeamMembers)
	v1.Post("/team/:projectId/:userId", pr.AddTeamMember)
	v1.Post("/kubeconfig", pr.GenerateKubeConfig)
}
