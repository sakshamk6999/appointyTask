package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	uri    = "mongodb+srv://tempUser:temporary@cluster0.bvxvh.mongodb.net/taskDB?retryWrites=true&w=majority"
	dbName = "taskDB"
)

func ConnectMongoDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("error while connecting to mongo : %v", err.Error())
	}

	return client, nil
}
