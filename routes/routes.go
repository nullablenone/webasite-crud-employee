package routes

import (
	"crud-employee/controllers"
	"net/http"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/test", controllers.Test)
}
