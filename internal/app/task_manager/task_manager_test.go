package taskmanager

import (
	"cli_tasks/internal/app/task"
	"testing"
)

var tm TaskManager

func TestGetTask(t *testing.T) {
	tm.Tasks = []task.Task{{Name: "read a book", Done: false}, {Name: "write a card", Done: false}}
	taskName := "write a card"
	task := tm.GetTask(taskName)

	if task == nil {
		t.Errorf("GetTask() error")
	}
}

func TestCreateTask(t *testing.T) {
	taskName := "do a thing"
	task := task.CreateTask(taskName)

	if task == nil {
		t.Errorf("CreateTask() error")
	}
}

func TestRemoveTask(t *testing.T) {
	tm.Tasks = []task.Task{{Name: "read a book", Done: false}, {Name: "write a card", Done: false}}
	taskName := "read a book"
	tm.RemoveTask(taskName)

	if tm.GetTask(taskName) != nil {
		t.Errorf("RemoveTask() error")
	}
}

func TestDoTask(t *testing.T) {
	tm.Tasks = []task.Task{{Name: "read a book", Done: false}, {Name: "write a card", Done: false}}
	taskName := "write a card"
	tm.DoTask(taskName)

	if task := tm.GetTask(taskName); task == nil || !task.Done {
		t.Errorf("DoTask() error")
	}
}

func TestAddTask(t *testing.T) {
	tm.Tasks = []task.Task{}
	taskName := "read a book"
	tm.AddTask(taskName)

	if task := tm.GetTask(taskName); task == nil {
		t.Errorf("AddTask() error")
	}
}
