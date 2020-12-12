package handlers

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"example.com/appointyTask/db"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMain(m *testing.M) {
	returnCode := m.Run()

	tearDown()

	os.Exit(returnCode)
}

func tearDown() {
	client, err := db.ConnectMongoDB()
	if err != nil {
		log.Fatal("error connecting to mongo")
	}

	database := client.Database("taskDB")
	collection := database.Collection("users")
	collection.DeleteMany(context.TODO(), bson.M{
		"id": "11",
	})
}

func TestUserHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserHandler)

	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	if status != 200 {
		t.Errorf("status code does not match !, got %v", status)
	}

}

func TestNoUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/4", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserHandler)

	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	if status != 400 {
		t.Errorf("status code does not match !, got %v", status)
	}
}

func TestCreateUser(t *testing.T) {
	var jsonStr = []byte(`
	{
		"id":"11",
		"name":"sak",
		"dateofBirth": "1999-09-06T00:00:00+00:00",
		"phoneNumber": "9999999988",
		"email": "exampe@gmail.com",
		"creationTime":"2020-12-12T00:00:00+00:00"
	}
	`)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserHandler)

	handler.ServeHTTP(responseRecorder, req)
	fmt.Println(responseRecorder.Body.String())
	status := responseRecorder.Code
	if status != 200 {
		t.Errorf("status code does not match !, got %v", status)
	}
}
