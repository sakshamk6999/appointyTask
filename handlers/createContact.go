package handlers

import (
	"context"
	"encoding/json"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client, err := db.ConnectMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database := client.Database("taskDB")
	collection := database.Collection("contact")

	result := collection.FindOne(context.TODO(), bson.M{
		"$and": []interface{}{
			bson.M{
				"userIdOne": insertContact.UserIdOne,
			},
			bson.M{
				"contacttime": insertContact.ContactTime,
			},
		},
	})
	err = result.Decode(&temp)
	if err != nil {
		http.Error(w, "user already has meeting at that time", http.StatusBadRequest)
		return
	}

	result = collection.FindOne(context.TODO(), bson.M{
		"$and": []interface{}{
			bson.M{
				"useridtwo": insertContact.UserIdOne,
			},
			bson.M{
				"contacttime": insertContact.ContactTime,
			},
		},
	})

	err = result.Decode(&temp)
	if err != nil {
		http.Error(w, "user already has meeting at that time", http.StatusBadRequest)
		return
	}

	insertResult, err := collection.InsertOne(context.TODO(), insertContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(insertResult)

}
