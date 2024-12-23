package routes

import (
	"my-go-web-server/src/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/hello/{username}", handlers.HelloHandler).Methods("GET")
	return router
}
