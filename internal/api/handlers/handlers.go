package handlers

import (
	taskm "cli_tasks/internal/app/task_manager"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
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

	var payload struct {
		TaskName string `json:"task_name"`
	}
	if err := json.Unmarshal(body, &payload); err != nil || payload.TaskName == "" {
		http.Error(w, "invalid task data", http.StatusBadRequest)
		return
	}

	tm.LoadTasks()
	tm.AddTask(payload.TaskName)
	tm.SaveTasks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func DoTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskName := strings.TrimPrefix(r.URL.Path, "/do/")
	if taskName == "" || taskName == r.URL.Path {
		http.Error(w, "Missing task_name parameter", http.StatusBadRequest)
		return
	}
	if unescaped, err := url.PathUnescape(taskName); err == nil {
		taskName = unescaped
	}

	tm.LoadTasks()
	tm.DoTask(taskName)
	tm.SaveTasks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskName := strings.TrimPrefix(r.URL.Path, "/remove/")
	if taskName == "" || taskName == r.URL.Path {
		http.Error(w, "Missing task_name parameter", http.StatusBadRequest)
		return
	}
	if unescaped, err := url.PathUnescape(taskName); err == nil {
		taskName = unescaped
	}

	tm.LoadTasks()
	tm.RemoveTask(taskName)
	tm.SaveTasks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tm.LoadTasks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tm.Tasks)
}
