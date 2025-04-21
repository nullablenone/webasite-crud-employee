package routes

import (
	"crud-employee/controllers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/employee", controllers.IndexEmployee(db)).Methods("GET")
	router.HandleFunc("/employee/create", controllers.CreateEmployee()).Methods("GET")
	router.HandleFunc("/employee/create", controllers.StoreEmployee(db)).Methods("POST")
	router.HandleFunc("/employee/edit/{id}", controllers.EditEmployee(db)).Methods("GET")
	router.HandleFunc("/employee/edit/{id}", controllers.MethodOverride(controllers.UpdateEmployee(db))).Methods("PUT", "POST")

	return router
}
