package main

import (
	"cli_tasks/cmd/api"
	"cli_tasks/cmd/cli"
	"cli_tasks/internal/database"
	"log"
	"os"
)

func main() {
	initStatus := make(chan bool, 2)
	go database.InitDatabase(initStatus)
	go api.InitApi(initStatus)

	if !<-initStatus {
		log.Fatal("Failed to initialize application components.")
		os.Exit(1)
	}

	cli.InitCLI()
}
