// Package controller Student API.
//
// the purpose of this application is to provide an application
// that is using go code to define an  Rest API
//
//     Schemes: http, https
//     Host: localhost:3000
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta

//go:generate swagger generate spec -m -o ./swagger.json

package main

import (
	"Univercity/handeller"
	"Univercity/utils"
	"context"
	"fmt"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	fmt.Println("Go  Tutorial")
	ctx, canc := context.WithTimeout(context.Background(), 10*time.Second)
	defer canc()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	utils.Client, _ = mongo.Connect(ctx, clientOptions)

	fmt.Println("connected")
	// Handle Subsequent requests

	defer utils.Client.Disconnect(ctx)
	handeller.HandleRequests()

}
