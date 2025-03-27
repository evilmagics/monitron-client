package main

import (
	"fmt"
	"log"
	"monitron-client/internal/database"
	"monitron-client/internal/routes"
	"monitron-client/internal/stats"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var AppVersion = "dev"

func main() {
	preStart()

	app := routes.Listen()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	onShutdown(app)
}

func preStart() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	os.Setenv("APP_VERSION", AppVersion)

	go stats.StartCollectStats(500 * time.Millisecond)
}

func onShutdown(app *fiber.App) {
	fmt.Println("Gracefully shutting down...")

	app.Shutdown()
	database.Cache().Flush()
	stats.StopCollectStats()

	fmt.Println("Application was successful shutdown.")
}
