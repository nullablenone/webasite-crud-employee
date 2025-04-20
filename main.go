package main

import (
	"crud-employee/config"
	"crud-employee/models"
	"crud-employee/routes"
	"net/http"
)

func main() {
	// database
	db := config.InitDataBase()
	err := models.MigrateEmployee(db)
	if err != nil {
		panic(err)
	}

	// route
	router := routes.SetupRoutes(db)
	http.ListenAndServe(":9090", router)
}
