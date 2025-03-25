package routes

import (
	"monitron-client/utils"

	"github.com/gofiber/fiber/v2"
)

func HandleNotFound() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Route Not Found",
		})

		return nil
	}
}

func HandleError() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		// Handle 404 Not Found error separately from other errors.
		// if e, ok := err.(*fiber.Error); ok && e == fiber.ErrNotFound {
		// 	c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		// 		"message": "Route Not Found",
		// 	})

		// 	return nil
		// }

		c.Status(fiber.StatusBadGateway).JSON(utils.FailedResponse(err))

		return nil
	}
}
