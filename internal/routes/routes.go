package routes

import "github.com/gofiber/fiber/v2"

func Routes() *fiber.App {
	app := app()

	Api(app)
	WebSocket(app)
	View(app)

	app.Use(HandleNotFound())

	return app
}
