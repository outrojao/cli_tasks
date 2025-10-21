package main

import (
	"cli_tasks/task"
	taskm "cli_tasks/task_manager"
	"testing"
)

var tm taskm.TaskManager

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
