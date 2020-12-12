package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"example.com/appointyTask/db"
	"example.com/appointyTask/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, err := db.ConnectMongoDB()
	if err != nil {
		log.Fatal("error connecting to mongo")
	}

	database := client.Database("taskDB")
	collection := database.Collection("users")

	if r.Method == "GET" {
		userId := r.URL.Path[len("/users/"):]
		filter := bson.M{
			"id": userId,
		}
		var resultUser models.User

		result := collection.FindOne(context.TODO(), filter)
		err = result.Decode(&resultUser)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(resultUser)
	} else if r.Method == "POST" {
		var insertUser models.User
		err := json.NewDecoder(r.Body).Decode(&insertUser)
		if err != nil {
			log.Fatal(err)
		}
		result, err := collection.InsertOne(context.TODO(), insertUser)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(result)
	}
}
