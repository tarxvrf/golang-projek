package controller

import "github.com/gofiber/fiber/v2"

func HandleControl(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ini kenapa dtang iwtku",
	})
}
