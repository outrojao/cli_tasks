package api

import (
	"bytes"
	"cli_tasks/api/routes"
	"cli_tasks/task"
	"encoding/json"
	"fmt"
	"net/http"
)

func InitApi(c chan bool) {
	routes.InitRoutes()
	go func() {
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			fmt.Println("Init server error:", err)
			c <- false
			return
		}
	}()
	c <- true
}

func CreateTask(url string, requiredStatusCode int, t *task.Task) {
	jsonBody, err := json.Marshal(t)
	if err != nil {
		fmt.Println("failed to marshal task:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != requiredStatusCode {
		fmt.Printf("Failed to create task. Status code: %d\n", resp.StatusCode)
	}
}
