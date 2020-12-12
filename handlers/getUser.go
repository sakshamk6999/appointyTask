package handlers

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"example.com/appointyTask/db"
// 	"example.com/appointyTask/models"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// func GetUserHandler(w http.ResponseWriter, r *http.Request) {
// 	userId := r.URL.Query().Get("userId")
// 	client, err := db.ConnectMongoDB()
// 	if err != nil {
// 		log.Fatal("error connecting to mongo")
// 	}

// 	database := client.Database("taskDB")
// 	collection := database.Collection("users")

// 	filter := bson.M{
// 		"id": userId,
// 	}
// 	var resultUser models.User

// 	result := collection.FindOne(context.TODO(), filter)
// 	err = result.Decode(&resultUser)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(resultUser)
// }
