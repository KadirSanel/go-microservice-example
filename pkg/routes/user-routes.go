package routes

import (
	"github.com/KadirSanel/go-microservice-example/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE")
}
