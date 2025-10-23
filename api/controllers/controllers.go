package controllers

import (
	"cli_tasks/task"
	taskm "cli_tasks/task_manager"
	"encoding/json"
	"net/http"
)

var tm *taskm.TaskManager = taskm.CreateTaskManager()

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CLI Task Manager API Endpoints\n"))
	w.Write([]byte("/create?task_name=YourTaskName\n"))
	w.Write([]byte("/do?task_name=YourTaskName"))
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTask = &task.Task{}

	contentType := r.Header.Get("Content-Type")
	if contentType == "application/json" || contentType == "application/json; charset=utf-8" {
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(newTask); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data: "+err.Error(), http.StatusBadRequest)
			return
		}
		name := r.FormValue("task_name")
		if name == "" {
			http.Error(w, "Missing task_name parameter", http.StatusBadRequest)
			return
		}
		newTask.Name = name
	}

	if newTask.Name == "" {
		http.Error(w, "task_name cannot be empty", http.StatusBadRequest)
		return
	}

	// convert task to json
	jsonData, err := json.MarshalIndent(newTask, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//read existing tasks from file
	tm.LoadTasks()

	// append new task to tasks slice
	tm.Tasks = append(tm.Tasks, *newTask)

	// save the new tasks into the file
	tm.SaveTasks()

	// respond with created task
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(jsonData)
}

func DoTask(w http.ResponseWriter, r *http.Request) {
	taskName := r.URL.Query().Get("task_name")
	if taskName == "" {
		http.Error(w, "Missing task_name parameter", http.StatusBadRequest)
		return
	}

	tm.LoadTasks()

	t := tm.GetTask(taskName)
	if t == nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	t.Do()

	jsonData, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tm.SaveTasks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonData)
}
