package main

import (
	"log"
	"net/http"

	"github.com/KadirSanel/go-microservice-example/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
