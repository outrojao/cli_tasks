package task

import "fmt"

type Task struct {
	//para interpretação da lib encoding/json é necessário que o nome dos atributos estejam em maisculos em refenciados em json type
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func (t *Task) Do() {
	if !t.Done {
		t.Done = true
		fmt.Printf("Task: '%s' is done!\n", t.Name)
	} else {
		fmt.Printf("Task: '%s' is already done!\n", t.Name)
	}
}

func CreateTask(taskName string) *Task {
	return &Task{
		taskName,
		false,
	}
}
