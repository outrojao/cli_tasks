package controllers

import (
	"cli_tasks/task"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/create?task_name=YourTaskName"))
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	// get task_name from query parameters
	taskName := r.URL.Query().Get("task_name")
	if taskName == "" {
		http.Error(w, "Missing task_name parameter", http.StatusBadRequest)
		return
	}

	// create new task
	newTask := task.CreateTask(taskName)

	// convert task to json
	jsonData, err := json.MarshalIndent(newTask, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// read existing tasks from file
	data, err := os.ReadFile("./tasks.json")
	if err != nil {
		fmt.Printf("Error to read file: %s", err.Error())
		http.Error(w, "Error to read tasks file", http.StatusInternalServerError)
		return
	}

	// unmarshal existing tasks
	var tasks []task.Task
	if len(data) > 0 {
		if err = json.Unmarshal(data, &tasks); err != nil {
			fmt.Printf("Could not parse tasks json file: %s", err.Error())
			http.Error(w, "Could not parse tasks json file", http.StatusInternalServerError)
			return
		}
	}

	// append new task to tasks slice
	tasks = append(tasks, *newTask)

	// marshal tasks back to json
	tasksData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Printf("Could not marshal tasks json: %s", err.Error())
		http.Error(w, "Could not marshal tasks json", http.StatusInternalServerError)
		return
	}

	// write updated tasks back to file
	if err = os.WriteFile("./tasks.json", tasksData, 0666); err != nil {
		fmt.Printf("Could not save tasks in json file: %s", err.Error())
		http.Error(w, "Could not save tasks in json file", http.StatusInternalServerError)
		return
	}

	// respond with created task
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(jsonData)
}
