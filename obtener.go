package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx := context.TODO()
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://Prueba:Welcome1@prueba1.bjtbaxh.mongodb.net/")
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	usersCollection := client.Database("Prueba").Collection("Usuarios")

	filter := bson.D{{"fullName", "Adolfo"}}

	// retrieve all the documents that match the filter
	cursor, err := usersCollection.Find(context.TODO(), filter)
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results from the search query")
	for _, result := range results {
		fmt.Println(result)

	}
}
