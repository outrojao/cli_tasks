package routes

import (
	"cli_tasks/api/controllers"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create", controllers.CreateTask)
	http.HandleFunc("/do", controllers.DoTask)
}
