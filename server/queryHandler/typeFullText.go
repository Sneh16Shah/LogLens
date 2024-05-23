package queryHandler

import (
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetFullTextFilter(queryParams url.Values) bson.M {
	fullText := queryParams.Get("textQuery")
	filter := bson.M{}
	orFilters := []bson.M{}
	textFilter := bson.M{"$regex": primitive.Regex{Pattern: fullText, Options: "i"}} // Case insensitive search

	// Define fields for text search
	textSearchFields := []string{"level", "message", "resourceId", "traceId", "spanId", "commit", "metadata.parentResourceId"}

	// Create $or filter for text search across specified fields
	for _, field := range textSearchFields {
		orFilters = append(orFilters, bson.M{field: textFilter})
	}

	filter["$or"] = orFilters
	return filter
}
