package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/metalstormbass/microservice-go/cmd/app"
	db "github.com/metalstormbass/microservice-go/cmd/database"
)


func main() {
	// DB Connection
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("An attempt was made... to connect to : %s", err.Error())
	}

	// App
	app := &app.App{
			Router: mux.NewRouter().StrictSlash(true),
			Database: database,
	}

	app.SetupRouter()
	

	log.Fatal(http.ListenAndServe(":8118", app.Router))
}