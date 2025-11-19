package main

import (
	"cli_tasks/internal/cli"
	"cli_tasks/internal/client"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Println("Error loading .env file:", err)
		return
	}

	if !client.HealthCheck() {
		log.Fatal("API health check failed")
		os.Exit(1)
	}

	cli.InitMenu()
}
