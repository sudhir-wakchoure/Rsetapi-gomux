package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, _ := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return client
	// }

	//fmt.Println("connection succesfully!", client)
	return client
}

var Client *mongo.Client

// Collection student

func NewFunction() *mongo.Collection {
	collection := Client.Database("univercity").Collection("student")
	return collection
}
