package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var mongo_uri = "mongodb://localhost:27017"

func Dbconnect() error {
	clientOptions := options.Client().ApplyURI(mongo_uri)
	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	return nil
}
