package main

import (
	"crud-employee/database"
	"crud-employee/routes"
	"net/http"
)

func main() {
	// database
	database.InitDataBase()

	// route
	server := http.NewServeMux()
	routes.MapRoutes(server)
	http.ListenAndServe(":9090", server)
}
