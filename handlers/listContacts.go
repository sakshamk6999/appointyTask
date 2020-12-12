package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/appointyTask/db"
	"example.com/appointyTask/models"
	"go.mongodb.org/mongo-driver/bson"
)

func listContactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, err := db.ConnectMongoDB()
	if err != nil {
		log.Fatal("error connecting to mongo")
	}

	database := client.Database("taskDB")
	collection := database.Collection("contact")
	userCollection := database.Collection("users")
	i := strings.Index(r.URL.Path, "&")
	userId := r.URL.Path[len("/contacts?user="):i]
	infectionTimeString := r.URL.Path[i+1:]

	tempInfectionTime, err := strconv.ParseInt(infectionTimeString, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	infectionTime := time.Unix(tempInfectionTime, 0)
	beforeTime := infectionTime.AddDate(0, 0, -14)

	cursor, err := collection.Find(context.TODO(), bson.M{
		"$and": []interface{}{
			bson.M{
				"userIdOne": userId,
			},
			bson.M{
				"contactTime": bson.M{
					"lte": infectionTime,
				},
			},
			bson.M{
				"contactTime": bson.M{
					"gte": beforeTime,
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
			log.Fatal(err)
		}

		filter := bson.M{
			"id": tempContact.UserIdTwo,
		}

		tempResult := userCollection.FindOne(context.TODO(), filter)
		err = tempResult.Decode(&tempUser)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, tempUser)
	}

	json.NewEncoder(w).Encode(result)
}