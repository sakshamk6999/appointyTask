package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"example.com/appointyTask/db"
	"example.com/appointyTask/models"
)

func createContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var insertContact models.Contact
	err := json.NewDecoder(r.Body).Decode(&insertContact)

	if err != nil {
		log.Fatal(err)
	}

	client, err := db.ConnectMongoDB()
	if err != nil {
		log.Fatal("error connecting to mongo")
	}

	database := client.Database("taskDB")
	collection := database.Collection("contact")

	result, err := collection.InsertOne(context.TODO(), insertContact)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)

}
