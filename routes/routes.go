package routes

import (
	"SIMPLE-API/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("Put")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	return router
}
