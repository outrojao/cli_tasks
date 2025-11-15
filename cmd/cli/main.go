package main

import (
	"cli_tasks/internal/api/client"
	"cli_tasks/internal/cli"
	"cli_tasks/internal/database"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatal("Error loading .env file:", err)
		os.Exit(1)
	}

	go database.InitDatabase(make(chan<- bool))
	defer database.CloseDatabase()

	if !client.HealthCheck() {
		log.Fatal("API health check failed")
		os.Exit(1)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-shutdown
		log.Println("Shutdown signal received")
		if err := database.CloseDatabase(); err != nil {
			log.Println("Error closing DB on shutdown:", err)
		}
		os.Exit(0)
	}()

	cli.InitMenu()
}
