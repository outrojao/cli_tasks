package api

import (
	"bytes"
	"cli_tasks/api/routes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

func InitApi(status chan<- bool) {
	routes.InitRoutes()

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Printf("Failed to start API server: %v\n", err)
		status <- false
		return
	}

	status <- true

	if err := http.Serve(ln, nil); err != nil {
		log.Printf("Failed to start HTTP server: %v\n", err)
	}
}

func CreateTask(taskName string) {
	jsonBody, err := json.Marshal(taskName)
	if err != nil {
		fmt.Println("failed to marshal task:", err)
		return
	}

	resp, err := http.Post("http://localhost:8000/create", "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Failed to create task. Status code: %d\n", resp.StatusCode)
	}
}

func DoTask(taskName string) {
	url := fmt.Sprintf("http://localhost:8000/do?task_name=%s", taskName)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to do task. Status code: %d\n", resp.StatusCode)
	}
}
