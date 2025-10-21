package main

import (
	"cli_tasks/task"
	taskm "cli_tasks/task_manager"
	"fmt"
	"os"

	"github.com/outrojao/mods/utils"
)

func main() {
	// apiStatus := make(chan bool)
	// go api.InitApi(apiStatus)
	// defer close(apiStatus)

	// apiOn := <-apiStatus
	// if !apiOn {
	// 	os.Exit(1)
	// }

	fmt.Printf("CLI - Task Manager\n\n")
	tm := taskm.CreateTaskManager("./tasks.json")
	tm.LoadTasks()
	for {
		utils.CreateMenu([]string{"Create task", "Do a task", "Remove a task", "List all tasks", "Load existing tasks", "Exit"}, "Menu")
		fmt.Println()
		switch option := utils.GetUserInput[int]("Choose a option: "); option {
		case 1:
			taskName := utils.GetUserInput[string]("Type the task name: ")
			tm.Tasks = append(tm.Tasks, *task.CreateTask(taskName))
		case 2:
			taskName := utils.GetUserInput[string]("Type the task name: ")
			task := tm.GetTask(taskName)
			task.Do()
		case 3:
			taskName := utils.GetUserInput[string]("Type the task name: ")
			tm.RemoveTask(taskName)
		case 4:
			tm.ListTasks()
		case 5:
			tm.LoadTasks()
		case 6:
			tm.SaveTasks()
			os.Exit(0)
		}
	}

}
