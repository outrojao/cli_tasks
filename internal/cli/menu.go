package cli

import (
	"cli_tasks/internal/api/client"
	"fmt"
	"strings"

	"github.com/outrojao/mods/utils"
)

func InitMenu() {
	fmt.Printf("CLI - Task Manager\n\n")
	for {
		utils.CreateMenu([]string{"Create task", "Do a task", "Remove a task", "List all tasks", "Exit"}, "Menu")
		fmt.Println()
		switch option := utils.GetUserInput[int]("Choose a option: "); option {
		case 1:
			taskName := strings.ToUpper(strings.TrimSpace(utils.GetUserInput[string]("Type the task name: ")))
			client.CreateTask(taskName)
		case 2:
			taskName := strings.ToUpper(strings.TrimSpace(utils.GetUserInput[string]("Type the task name: ")))
			client.DoTask(taskName)
		case 3:
			taskName := strings.ToUpper(strings.TrimSpace(utils.GetUserInput[string]("Type the task name: ")))
			client.RemoveTask(taskName)
		case 4:
			client.ListTasks()
		case 5:
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}

}
