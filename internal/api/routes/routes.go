package routes

import (
	"cli_tasks/internal/api/handlers"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/create", handlers.CreateTask)
	http.HandleFunc("/do/", handlers.DoTask)
	http.HandleFunc("/remove/", handlers.RemoveTask)
	http.HandleFunc("/list", handlers.ListTasks)
}
