package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/metalstormbass/microservice-go/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB Connection String
const db_host = "goofydbservice"
const db_port = "27017"
const db_collection = "goofycollection"
var collection *mongo.Collection


// Special init function to set up DB connection
func init()  {

	// Build Connection Strings
	connect_string := fmt.Sprintf("mongodb://%s:%s", db_host, db_port)
	
	clientOptions := options.Client().ApplyURI(connect_string)
	
	// Connect to DB
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
collection = client.Database(db_host).Collection(db_collection)
	log.Print("YESSSSS. It worked. Connected to MongoDB")

}

// Routing

// GET All Tasks Function
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)
	log.Println("You reached into that database and ripped out all the things")

}

// POST function to Create a TASK
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.DbData
	_ = json.NewDecoder(r.Body).Decode(&task)

	insertTask(task)
	json.NewEncoder(w).Encode(task)
	log.Println("You called a thinggggggg")
	w.WriteHeader(http.StatusOK)
}


// CRUD Functions
// Get all tasks
func getAllTasks() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	

	if err != nil {
		log.Print(err.Error())
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Print(e)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Print(err.Error())
	}
	
	cur.Close(context.Background())
	return results
}


// Create a task 
func insertTask(task models.DbData) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Print(err.Error())
	}
	log.Println("Added Task: ", insertResult.InsertedID)
}