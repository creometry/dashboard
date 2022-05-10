package controllers

import (
	"github.com/Creometry/dashboard/resource/project"
	"github.com/gofiber/fiber/v2"
)


func CreateVCluster(c *fiber.Ctx)error{
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

	kubeconfig,err:=project.CreateVCluster(reqData.Plan, reqData.UsrProjectName)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"kubeconfig":kubeconfig,
	})
}