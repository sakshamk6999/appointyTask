package models

import "time"

type Contact struct {
	UserIdOne   string    `json:"userIdOne" bson:"userIdOne"`
	UserIdTwo   string    `json: "userIdTwo" bson: "userIdTwo"`
	ContactTime time.Time `json: "contactTime" bson: "contactTime"`
}
