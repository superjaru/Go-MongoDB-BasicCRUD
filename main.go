package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
    // Connect to MongoDB
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    var err error
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

    collection := client.Database("testdb").Collection("testcollection")

    // Insert a document
    newDoc := bson.D{
        {Key: "name", Value: "John Doe"},
        {Key: "age", Value: 30},
    }
    insertResult, err := collection.InsertOne(context.TODO(), newDoc)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Inserted document: ", insertResult.InsertedID)

    // Find a document
    var result bson.D
    err = collection.FindOne(context.TODO(), bson.D{{Key: "name", Value: "John Doe"}}).Decode(&result)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Found document: ", result)

    // Update a document
    filter := bson.D{
        {Key: "name", Value: "John Doe"},
    }
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "age", Value: 31}}}}

    updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

    // Delete a document
    deleteResult, err := collection.DeleteOne(context.TODO(), filter)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Deleted %v documents in the testcollection collection\n", deleteResult.DeletedCount)

    // Close the connection
    err = client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connection to MongoDB closed.")
}
