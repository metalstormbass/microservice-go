package router

import (
	"github.com/gorilla/mux"
	"github.com/metalstormbass/microservice-go/src/app"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Routes Traffic to appropriate app function
	router.HandleFunc("/api/task", app.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", app.CreateTask).Methods("POST", "OPTIONS")

	return router
}