package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router *mux.Router
	Database *mongo.Client
}


// Setup Router
func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Path("/endpoint/{id}").
		HandlerFunc(app.getFunction)
	
	app.Router.
		Methods("POST").
		Path("/endpoint").
		HandlerFunc(app.postFunction)
}

// GET function
func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	
	log.Println("You reached into that database and ripped out a thing")
	w.WriteHeader(http.StatusOK)

}


// POST function
func (app *App) postFunction(w http.ResponseWriter, r *http.Request) {
	log.Println("You called a thinggggggg")
	w.WriteHeader(http.StatusOK)
}