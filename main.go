package main

import (
	"cli_tasks/api"
	taskm "cli_tasks/task_manager"
	"fmt"
	"log"
	"os"

	"github.com/outrojao/mods/utils"
)

func main() {
	apiStatus := make(chan bool, 1)
	go api.InitApi(apiStatus)

	if ok := <-apiStatus; !ok {
		log.Fatal("Failed to start API server")
		os.Exit(1)
	}

	tm := taskm.CreateTaskManager()
	fmt.Printf("CLI - Task Manager\n\n")
	for {
		tm.LoadTasks()
		utils.CreateMenu([]string{"Create task", "Do a task", "Remove a task", "List all tasks", "Load existing tasks", "Exit"}, "Menu")
		fmt.Println()
		switch option := utils.GetUserInput[int]("Choose a option: "); option {
		case 1:
			taskName := utils.GetUserInput[string]("Type the task name: ")
			api.CreateTask(taskName)
		case 2:
			taskName := utils.GetUserInput[string]("Type the task name: ")
			api.DoTask(taskName)
		case 3:
			taskName := utils.GetUserInput[string]("Type the task name: ")
			api.RemoveTask(taskName)
		case 4:
			tm.ListTasks()
		case 5:
			tm.LoadTasks()
		case 6:
			os.Exit(0)
		}
	}

}
