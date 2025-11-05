package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCreateTaskApi(t *testing.T) {
	payload := `{"task_name":"TestTask"}`
	req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	CreateTask(rr, req)
	defer tm.RemoveTask("TestTask")

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusCreated, rr.Code, rr.Body.String())
	}
}

func TestDoTaskApi(t *testing.T) {
	taskName := "TestTask"
	tm.AddTask(taskName)
	defer tm.RemoveTask(taskName)

	taskNameEscaped := url.PathEscape(taskName)
	url := "/do/" + taskNameEscaped
	req := httptest.NewRequest(http.MethodPut, url, nil)
	rr := httptest.NewRecorder()

	DoTask(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusOK, rr.Code, rr.Body.String())
	}
}

func TestRemoveTaskApi(t *testing.T) {
	taskName := "TestTask"
	tm.AddTask(taskName)

	taskNameEscaped := url.PathEscape(taskName)
	url := "/remove/" + taskNameEscaped
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	rr := httptest.NewRecorder()

	RemoveTask(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusOK, rr.Code, rr.Body.String())
	}
}

func TestListTasksApi(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	rr := httptest.NewRecorder()

	ListTasks(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusOK, rr.Code, rr.Body.String())
	}
}
