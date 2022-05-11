package user

import (
	"github.com/Creometry/dashboard/resource/user"
	"github.com/gofiber/fiber/v2"
)


func AddUserToProject(c *fiber.Ctx) error {
	var reqData user.ReqData
	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := reqData.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	data,err:=user.AddUserToProject(reqData)

	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message":"user added to project",
		"data":data,
	})


}