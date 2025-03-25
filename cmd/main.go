package main

import (
	"fmt"
	"log"
	"monitron-client/internal/routes"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

var AppVersion = "dev"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	os.Setenv("APP_VERSION", AppVersion)

	app := routes.Listen()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Fiber was successful shutdown.")
}
