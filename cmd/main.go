package main

import (
	"cli_tasks/internal/cli"
	"cli_tasks/internal/client"
	"log"
	"os"
)

func main() {
	if !client.HealthCheck() {
		log.Fatal("API health check failed")
		os.Exit(1)
	}

	cli.InitMenu()
}
