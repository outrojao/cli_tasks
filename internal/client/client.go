package client

import (
	"bytes"
	"cli_tasks/internal/app/task"
	"cli_tasks/internal/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

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
	jwtToken, err := auth.CreateToken()
	if err != nil {
		fmt.Println("failed to create JWT token:", err)
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+jwtToken)

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

	jwtToken, err := auth.CreateToken()
	if err != nil {
		fmt.Println("failed to create JWT token:", err)
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+jwtToken)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		fmt.Printf("Task is already done!\n")
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

	jwtToken, err := auth.CreateToken()
	if err != nil {
		fmt.Println("failed to create JWT token:", err)
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+jwtToken)

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

	jwtToken, err := auth.CreateToken()
	if err != nil {
		fmt.Println("failed to create JWT token:", err)
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+jwtToken)

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

func HealthCheck() bool {
	url := "http://localhost:8000/health"
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("request error:", err)
		return false
	}
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		fmt.Println("request error:", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
