package controllers

import (
	taskm "cli_tasks/task_manager"
	"encoding/json"
	"io"
	"net/http"
)

var tm *taskm.TaskManager = taskm.CreateTaskManager()

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CLI Task Manager API Endpoints\nPOST - /create\nGET - /do?task_name=YourTaskName"))
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	var plain string
	if err := json.Unmarshal(body, &plain); err != nil || plain == "" {
		http.Error(w, "invalid task data", http.StatusBadRequest)
		return
	}

	tm.LoadTasks()
	tm.AddTask(plain)
	tm.SaveTasks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func DoTask(w http.ResponseWriter, r *http.Request) {
	taskName := r.URL.Query().Get("task_name")
	if taskName == "" {
		http.Error(w, "Missing task_name parameter", http.StatusBadRequest)
		return
	}

	tm.LoadTasks()
	tm.DoTask(taskName)
	tm.SaveTasks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
