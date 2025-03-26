package routes

import (
	"log"
	"monitron-client/internal/stats"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Routes() *fiber.App {
	app := app()

	stat := app.Group("/api/stats")
	{
		stat.Get("/", stats.HandleAPI)
		stat.Get("/cpu", stats.HandleCPU)
		stat.Get("/cpu/info", stats.HandleCPU)
		stat.Get("/host", stats.HandleHost)
		stat.Get("/disk", stats.HandleDisk)
		stat.Get("/disk/partition", stats.HandleDisk)
		stat.Get("/memory", stats.HandleMemory)
		stat.Get("/network", stats.HandleNetwork)
	}

	ws := app.Group("/ws/stats")
	{
		ws.Get("/cpu", stats.HandleCPUWS())
		ws.Get("/memory", stats.HandleMemoryWS())
		ws.Get("/disk", stats.HandleDiskWS())
		ws.Get("/network", stats.HandleNetworkWS())
	}
	app.Use(HandleNotFound())
	return app
}

func Listen() *fiber.App {
	appUrl, ok := os.LookupEnv("APP_URL")
	if !ok {
		appUrl = "0.0.0.0:9898"
	}

	app := Routes()
	go func() {
		if err := app.Listen(appUrl); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	return app
}

func app() *fiber.App {
	version := os.Getenv("APP_VERSION")
	app := fiber.New(fiber.Config{
		AppName:      "Monitron Client " + version,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  1 * time.Hour,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ServerHeader: "PsUtils",
		ErrorHandler: HandleError(),
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(compress.New())

	return app
}
