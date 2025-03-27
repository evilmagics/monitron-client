package routes

import (
	"log"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Listen() *fiber.App {
	appURL, ok := os.LookupEnv("APP_URL")
	if !ok {
		appURL = "0.0.0.0:9898"
	}

	app := Routes()

	// Handle
	go func() {
		if err := app.Listen(appURL); err != nil {
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
