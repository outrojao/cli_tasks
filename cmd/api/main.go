package api

import (
	"bytes"
	"cli_tasks/internal/api/routes"
	"cli_tasks/internal/app/task"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
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
	payload := map[string]string{"task_name": taskName}
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("failed to marshal task:", err)
		return
	}

	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8000/create", bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(httpReq)
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
	url := fmt.Sprintf("http://localhost:8000/do/%s", url.PathEscape(taskName))
	httpReq, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf("Task is already done. Status code: %d\n", resp.StatusCode)
	} else if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to do task. Status code: %d\n", resp.StatusCode)
	}
}

func RemoveTask(taskName string) {
	url := fmt.Sprintf("http://localhost:8000/remove/%s", url.PathEscape(taskName))
	httpReq, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to remove task. Status code: %d\n", resp.StatusCode)
	}
}

func ListTasks() {
	url := "http://localhost:8000/list"
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to list tasks. Status code: %d\n", resp.StatusCode)
		return
	}

	var tasks []task.Task
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		fmt.Println("failed to decode response:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("You don't have any task!")
		return
	}

	fmt.Println("Yours tasks:")
	for i, task := range tasks {
		if task.Done {
			fmt.Printf("#%d - %s (done)\n", i+1, task.Name)
			continue
		}
		fmt.Printf("#%d - %s\n", i+1, task.Name)
	}
	fmt.Println()
}
