package controllers

import (
	"fmt"
	"os"

	"github.com/Creometry/dashboard/go-provisioner/internal/project"
	"github.com/Creometry/dashboard/go-provisioner/internal/team"
	"github.com/gofiber/fiber/v2"
)

func ProvisionProject(c *fiber.Ctx) error {
	// parse the request body
	reqData := new(project.ReqData)
	if err := c.BodyParser(reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// check if the request body is valid
	if err := reqData.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	data, err := project.ProvisionProject(*reqData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"projectId": data.ProjectId,
	})
}

func ProvisionProjectNewUser(c *fiber.Ctx) error {
	// parse the request body
	reqData := new(project.ReqDataNewUser)
	if err := c.BodyParser(reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// check if the request body is valid
	if err := reqData.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	data, err := project.ProvisionProjectNewUser(*reqData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"projectId": data.ProjectId,
		"token":     data.Token,
		"password":  data.Password,
	})
}

func GenerateKubeConfig(c *fiber.Ctx) error {
	// get the token from the body
	reqData := new(project.ReqDataKubeconfig)
	if err := c.BodyParser(reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// check if the request body is valid
	if reqData.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "token is required",
		})
	}
	data, err := project.GetKubeConfig(reqData.Token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"config": data,
	})
}

// func FindUserAndLoginOrCreate(c *fiber.Ctx) error {
// 	// get username from path
// 	username := c.Params("username")
// 	if username == "" {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "username is required",
// 		})
// 	}

// 	data, err := project.FindUser(username)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"user_token": data.Token,
// 		"user_id":    data.Id,
// 		"namespace":  data.Namespace,
// 		"projectId":  data.ProjectId,
// 	})
// }

func ListTeamMembers(c *fiber.Ctx) error {
	projectId := c.Params("projectId")
	if projectId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "projectId is required",
		})
	}
	data, err := team.ListTeamMembers(fmt.Sprintf("%s:%s", projectId, os.Getenv("CLUSTER_ID")))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"members": data,
	})
}

func AddTeamMember(c *fiber.Ctx) error {
	projectId := c.Params("projectId")
	username := c.Params("userId")

	if projectId == "" || username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "projectId and username are required",
		})
	}
	userId, _, err := project.GetUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	data, err := project.AddUserToProject(userId, projectId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

func Login(c *fiber.Ctx) error {
	// get the token from the body
	reqData := new(project.ReqDataLogin)
	if err := c.BodyParser(reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// check if the request body is valid
	if err := reqData.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	token, err := project.Login(reqData.Username, reqData.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"token": token,
	})
}
