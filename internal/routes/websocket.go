package routes

import (
	"monitron-client/internal/stats"

	"github.com/gofiber/fiber/v2"
)

func WebSocket(app *fiber.App) {
	ws := app.Group("/ws/stats")
	{
		ws.Get("/cpu/usage", stats.HandleCPUUsageWS())
		ws.Get("/memory/usage", stats.HandleMemoryUsageWS())
		ws.Get("/disks/usage", stats.HandleDiskUsageWS())
		ws.Get("/network/usage", stats.HandleNetworkUsageWS())
	}
}
