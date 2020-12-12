package models

import "time"

type User struct {
	ID           string    `json:"id" bson:"id"`
	Name         string    `json: "name" bson: "name"`
	DateOfBirth  time.Time `json: "dateOfBirth" bson: "dateOfBirth"`
	PhoneNumber  string    `json: "phone" bson: "phone"`
	Email        string    `json: "email" bson: "email"`
	CreationTime time.Time `json: "creationTime" bson: "creationTime"`
}
