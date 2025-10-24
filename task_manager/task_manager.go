package taskmanager

import (
	"cli_tasks/task"
	"encoding/json"
	"fmt"
	"os"
)

type TaskManager struct {
	localJsonDir string
	Tasks        []task.Task
}

func (tm TaskManager) GetTask(taskName string) *task.Task {
	for i := range tm.Tasks {
		if tm.Tasks[i].Name == taskName {
			return &tm.Tasks[i] // ← Retorna ponteiro para o elemento original
		}
	}
	return nil
}

func (tm *TaskManager) AddTask(taskName string) *task.Task {
	newTask := task.CreateTask(taskName)
	tm.Tasks = append(tm.Tasks, *newTask)
	return newTask
}

func (tm *TaskManager) RemoveTask(taskName string) {
	for i := range tm.Tasks {
		if tm.Tasks[i].Name == taskName {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			break
		}
	}
}

func (tm *TaskManager) DoTask(taskName string) {
	t := tm.GetTask(taskName)
	if t != nil {
		t.Do()
	}
}

func (tm TaskManager) ListTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Println("You don't have any task!")
		return
	}

	fmt.Println("Yours tasks:")
	for i, task := range tm.Tasks {
		fmt.Printf("#%d - %s\n", i+1, task.Name)
	}
	fmt.Println()
}

func (tm *TaskManager) SaveTasks() {
	jsonData, err := json.MarshalIndent(tm.Tasks, "", " ")
	if err != nil {
		fmt.Printf("Could not convert struct data into json file: %s\n", err.Error())
		return
	}

	if err = os.WriteFile(tm.localJsonDir, jsonData, 0644); err != nil {
		fmt.Printf("Could not save tasks in json file: %s\n", err.Error())
	}
}
func (tm *TaskManager) LoadTasks() {
	data, err := os.ReadFile(tm.localJsonDir)
	if err != nil {

		if os.IsNotExist(err) {
			tm.Tasks = []task.Task{}
			return
		}
		fmt.Printf("Error reading file: %s\n", err.Error())
		return
	}

	err = json.Unmarshal(data, &tm.Tasks)
	if err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err.Error())
		return
	}
}

func CreateTaskManager() *TaskManager {
	return &TaskManager{
		localJsonDir: "./tasks.json",
		Tasks:        []task.Task{},
	}
}
