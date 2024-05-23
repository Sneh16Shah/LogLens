package routeHandler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sneh16Shah/dytelogingestor/server/utils"
	"github.com/Sneh16Shah/dytelogingestor/server/queryHandler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func QueryLogsHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil) 	//check the connection
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(utils.Database).Collection(utils.Collection)	// Get a handle for your MongoDB collection
	
	//construct filter
	filter := bson.M{} 
	switch {
	case queryParams.Get("textQuery") != "":
		filter = queryHandler.GetFullTextFilter(queryParams)
	case queryParams.Get("startDate")!="" || queryParams.Get("endDate")!="":
		filter = queryHandler.GetTimeRangeFilter(queryParams)
	default:
		filter = queryHandler.GetMuiltiFilter(queryParams)
 	}

	cursor, err := collection.Find(context.Background(), filter) 	// Query MongoDB with the constructed filter
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var logs []utils.Log
	for cursor.Next(context.Background()) {
		var logData utils.Log
		err := cursor.Decode(&logData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		logs = append(logs, logData)
	}

	logsJSON, err := json.Marshal(logs) 	// Convert logs to JSON and send as the response
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(logsJSON)
}
