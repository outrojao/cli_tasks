package task

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	taskName := "do a thing"
	task := CreateTask(taskName)

	if task == nil {
		t.Errorf("CreateTask() error")
	}
}

func TestDoTask(t *testing.T) {
	taskName := "write a card"
	task := CreateTask(taskName)
	task.Do()
	if !task.Done {
		t.Errorf("Do() error")
	}
}
