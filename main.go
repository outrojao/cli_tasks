package main

import (
	"cli_tasks/api"
	taskm "cli_tasks/task_manager"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

	fmt.Printf("CLI - Task Manager\n\n")
	tm := taskm.CreateTaskManager()
	tm.LoadTasks()
	for {
		utils.CreateMenu([]string{"Create task", "Do a task", "Remove a task", "List all tasks", "Load existing tasks", "Exit"}, "Menu")
		fmt.Println()
		switch option := utils.GetUserInput[int]("Choose a option: "); option {
		case 1:
			taskName := utils.GetUserInput[string]("Type the task name: ")
			url := fmt.Sprintf("http://localhost:8000/create?task_name=%s", url.QueryEscape(taskName))
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			// read and discard body to ensure full response transmission, then close
			_, _ = io.ReadAll(resp.Body)
			resp.Body.Close()

			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("Failed to create task. Status code: %d\n", resp.StatusCode)
				continue
			}

			tm.LoadTasks()
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
