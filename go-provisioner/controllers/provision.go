package controllers

import (
	"fmt"
	"strings"

	"github.com/Creometry/dashboard/go-provisioner/internal/project"
	"github.com/Creometry/dashboard/go-provisioner/internal/team"
	"github.com/Creometry/dashboard/go-provisioner/utils"
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

func ListTeamMembers(c *fiber.Ctx) error {
	projectId := c.Params("projectId")
	if projectId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "projectId is required",
		})
	}
	// if projectId contains ':' then list team members without change
	var prId string
	if strings.Contains(projectId, ":") {
		prId = projectId
	} else {
		clusterId, err := utils.GetVariable("config", "CLUSTER_ID")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		prId = fmt.Sprintf("%s:%s", clusterId, projectId)
	}
	data, err := team.ListTeamMembers(prId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if data == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no members found or invalid projectId",
		})
	}
	return c.JSON(fiber.Map{
		"members": data,
		"prId":    prId,
	})
}

func AddTeamMember(c *fiber.Ctx) error {
	projectId := c.Params("projectId")
	userId := c.Params("userId")

	if projectId == "" || userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "projectId and userId are required",
		})
	}

	var prId string
	if strings.Contains(projectId, ":") {
		prId = projectId
	} else {
		clusterId, err := utils.GetVariable("config", "CLUSTER_ID")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		prId = fmt.Sprintf("%s:%s", clusterId, projectId)
	}

	data, err := project.AddUserToProject(userId, prId)
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
	token, id, uuid, err := project.Login(reqData.Username, reqData.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"token":  token,
		"userId": id,
		"uuid":   uuid,
	})
}

func Register(c *fiber.Ctx) error {
	// get the token from the body
	reqData := new(project.ReqDataRegister)
	if err := c.BodyParser(reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// check if the request body is valid
	if reqData.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "username is required",
		})
	}
	id, token, password, uuid, err := project.Register(reqData.Username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if token == "" || id == "" || password == "" || uuid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error registering user",
		})
	}

	return c.JSON(fiber.Map{
		"token":    token,
		"userId":   id,
		"password": password,
		"uuid":     uuid,
	})
}
