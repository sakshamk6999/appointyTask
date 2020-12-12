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

func CreateContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var insertContact models.Contact
	var temp models.Contact
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

	result := collection.FindOne(context.TODO(), bson.M{
		"$and": []interface{}{
			bson.M{
				"userIdOne": insertContact.UserIdOne,
			},
			bson.M{
				"contactTime": insertContact.ContactTime,
			},
		},
	})
	err = result.Decode(&temp)
	if err != nil {
		log.Fatal("user ", insertContact.UserIdOne, "already has a meeting at time ", insertContact.ContactTime)
	}

	result = collection.FindOne(context.TODO(), bson.M{
		"$and": []interface{}{
			bson.M{
				"userIdTwo": insertContact.UserIdOne,
			},
			bson.M{
				"contactTime": insertContact.ContactTime,
			},
		},
	})

	err = result.Decode(&temp)
	if err != nil {
		log.Fatal("user ", insertContact.UserIdTwo, " already has a meeting at time ", insertContact.ContactTime)
	}

	insertResult, err := collection.InsertOne(context.TODO(), insertContact)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(insertResult)

}
