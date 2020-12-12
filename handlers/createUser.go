package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"example.com/appointyTask/db"
	"example.com/appointyTask/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var insertUser models.User

	err := json.NewDecoder(r.Body).Decode(&insertUser)

	if err != nil {
		log.Fatal(err)
	}

	client, err := db.ConnectMongoDB()
	if err != nil {
		log.Fatal("error connecting to mongo")
	}

	database := client.Database("taskDB")
	collection := database.Collection("users")

	result, err := collection.InsertOne(context.TODO(), insertUser)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}
