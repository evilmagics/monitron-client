package routes

import (
	"monitron-client/internal/stats"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	stat := app.Group("/api/stats")
	{
		stat.Get("/", stats.HandleStats)
		stat.Get("/host", stats.HandleHost)
		stat.Get("/cpu", stats.HandleCPU)
		stat.Get("/cpu/info", stats.HandleCPUInfo)
		stat.Get("/cpu/usage", stats.HandleCPUUsage)
		stat.Get("/disks", stats.HandleDisk)
		stat.Get("/disks/usage", stats.HandleDiskUsage)
		stat.Get("/disks/partitions", stats.HandleDiskPartitions)
		stat.Get("/memory", stats.HandleMemory)
		stat.Get("/memory/usage", stats.HandleMemoryUsage)
		stat.Get("/network/usage", stats.HandleNetworkUsage)
	}
}
