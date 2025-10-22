package api

import (
	"cli_tasks/api/routes"
	"fmt"
	"net/http"
)

func InitApi(c chan bool) {
	routes.InitRoutes()
	go func() {
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			fmt.Println("Init server error:", err)
			c <- false
			return
		}
	}()
	c <- true
}
