package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET")
		Path("/endpoint/{id}")
		HandlerFunc(app.getFunction)
	
	app.Router.
		Methods("POST")
		Path("/endpoint")
		HandlerFunc(app.postFunction)
}

// GET function
func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id, ok := vars["id"]
	if !ok {
		log.Fatal("FATALITY - ID not found in path")
	}
	
	dbdata := &DbData{}
	err := app.Database.QueryRow("SELECT id, fate, name FROM 'goofytable' where id = ?", id).Scan(&dbdata.ID, &dbdata.Data, &dbdata.Name)
	if err != nil {
		log.Fatal("FATALITY - SELECT failed")
	}

	log.Println("You reached into that database and ripped out a thing")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbdata); err != nill {
		panic(err)
	}

}


// POST function
func (app *App) postFunction(w http.ResponseWriter, r *http.Request) {
	_, err = database.Exec("INSERT INTO 'goofytable' (name) VALUES ('myname')")
	if err != nil {
		log.Fatal("An attempt was made... to INSERT")
	}

	log.Println("You called a thinggggggg")
	w.WriteHeader(http.StatusOK)
}