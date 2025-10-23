package main

import (
	"cli_tasks/api"
	"cli_tasks/task"
	taskm "cli_tasks/task_manager"
	"fmt"
	"net/http"
	"os"

	"github.com/outrojao/mods/utils"
)

func main() {
	apiStatus := make(chan bool)
	go api.InitApi(apiStatus)
	defer close(apiStatus)

	apiOn := <-apiStatus
	if !apiOn {
		os.Exit(1)
	}

	tm := taskm.CreateTaskManager()
	tm.LoadTasks()
	fmt.Printf("CLI - Task Manager\n\n")
	for {
		utils.CreateMenu([]string{"Create task", "Do a task", "Remove a task", "List all tasks", "Load existing tasks", "Exit"}, "Menu")
		fmt.Println()
		switch option := utils.GetUserInput[int]("Choose a option: "); option {
		case 1:
			taskName := utils.GetUserInput[string]("Type the task name: ")
			t := task.CreateTask(taskName)
			api.CreateTask("http://localhost:8000/create", http.StatusCreated, t)
		case 2:
			// taskName := utils.GetUserInput[string]("Type the task name: ")
			// url := fmt.Sprintf("http://localhost:8000/do?task_name=%s", url.QueryEscape(taskName))
			// api.FetchApi(url, http.StatusOK)
		case 3:
			taskName := utils.GetUserInput[string]("Type the task name: ")
			tm.RemoveTask(taskName)
		case 4:
			tm.ListTasks()
		case 5:
			tm.LoadTasks()
		case 6:
			os.Exit(0)
		}
	}

}
