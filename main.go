package main

import (
	"cli_tasks/cmd/api"
	"cli_tasks/cmd/cli"
	"cli_tasks/internal/database"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dotenv-org/godotenvvault"
)

func main() {
	if err := godotenvvault.Load("configs/.env"); err != nil {
		log.Fatal("Error loading .env file:", err)
		os.Exit(1)
	}

	initStatus := make(chan bool, 2)
	go database.InitDatabase(initStatus)
	defer database.CloseDatabase()

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

	go api.InitApi(initStatus)

	if !<-initStatus {
		log.Fatal("Failed to initialize application components.")
		os.Exit(1)
	}

	cli.InitCLI()
}
