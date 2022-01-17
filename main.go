package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/metalstormbass/microservice-go/cmd/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database Struct
type DbData struct {
	ID   int       `json:"id"`
	Date time.Time `json:"date"`
	Name string    `json:"name"`
}

func db () *mongo.Client {
	db_host := "goofydbservice"
	db_port := "27017"

	// Connect to Database
	connect_string := fmt.Sprintf("mongodb://%s:%s", db_host, db_port)
	fmt.Println(connect_string)
	clientOptions := options.Client().ApplyURI(connect_string)
	fmt.Println(clientOptions)
	
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Print("DB Connect Failed")
		log.Fatal(err.Error())
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
	log.Fatal(err)
}

	log.Print("YESSSSS. It worked. Connected to MongoDB")

	return client
}



func main() {
	// DB Connection + Instantiation
	database := db()
    
	//var goofyCollection = db().Database("goofydb").Collection("goofy")

	
	// App
	app := &app.App{
			Router: mux.NewRouter().StrictSlash(true),
			Database: database,
	}
	app.SetupRouter()
	
	log.Print(http.ListenAndServe(":8118", app.Router))
}