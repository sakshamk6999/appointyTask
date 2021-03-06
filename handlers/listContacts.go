package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/appointyTask/db"
	"example.com/appointyTask/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListContactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, err := db.ConnectMongoDB()
	if err != nil {
		http.Error(w, "error connecting to mongo", http.StatusInternalServerError)
		// log.Fatal("error connecting to mongo")
		return
	}

	database := client.Database("taskDB")
	collection := database.Collection("contact")
	userCollection := database.Collection("users")
	i := strings.Index(r.URL.String(), "&")
	urlString := r.URL.String()
	userId := urlString[len("/contacts?user="):i]
	infectionTimeString := urlString[i+1+len("infection_timestamp="):]

	tempInfectionTime, err := strconv.ParseInt(infectionTimeString, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	infectionTime := time.Unix(tempInfectionTime, 0)
	fmt.Println(infectionTime.UTC().String())
	beforeTime := infectionTime.AddDate(0, 0, -14)
	fmt.Println(beforeTime.UTC())
	cursor, err := collection.Find(context.TODO(), bson.M{
		"$and": []interface{}{
			bson.M{
				"userIdOne": userId,
			},
			bson.M{
				"contacttime": bson.M{
					"$lte": infectionTime.UTC(),
				},
			},
			bson.M{
				"contacttime": bson.M{
					"$gte": beforeTime.UTC(),
				},
			},
		},
	})

	var result []models.User

	for cursor.Next(context.Background()) {
		var tempContact models.Contact
		var tempUser models.User
		err := cursor.Decode(&tempContact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			// log.Fatal(err)
			return
		}

		filter := bson.M{
			"id": tempContact.UserIdTwo,
		}

		tempResult := userCollection.FindOne(context.TODO(), filter)
		err = tempResult.Decode(&tempUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			// log.Fatal(err)
			return
		}

		result = append(result, tempUser)
	}

	json.NewEncoder(w).Encode(result)
}
