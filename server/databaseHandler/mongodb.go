package databasehandler

import (
	"context"
	"log"

	"github.com/Sneh16Shah/dytelogingestor/server/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func SaveToMongoDB(newlog utils.Log) {
	collection := ConnectToMongoDB()
	_, err := collection.InsertOne(context.Background(), newlog)
	if err != nil {
		log.Println("Failed to insert log into MongoDB:", err)
		// Handle error
	}
}

func ConnectToMongoDB() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(utils.Database).Collection(utils.Collection)
	return collection
}

