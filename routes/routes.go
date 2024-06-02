package routes

import (
	"SIMPLE-API/controllers"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/users", controllers.UsersHandler)
	router.HandleFunc("/users/", controllers.UserHandler)

	return router
}
