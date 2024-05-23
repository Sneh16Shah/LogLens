package queryHandler

import (
	"encoding/json"
	"fmt"
	"net/url"
	
	"github.com/Sneh16Shah/dytelogingestor/server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetMuiltiFilter(queryParams url.Values) bson.M{
	filter := bson.M{}
	var andFilters []bson.M

	for key, value := range queryParams {
		if key == "metadata" {
			var metadata utils.Metadata
			if err := json.Unmarshal([]byte(value[0]), &metadata); err != nil {
				fmt.Println("Error:", err)
			}
			if metadata.ParentResourceId != "" {
				textFilter := bson.M{"$regex": primitive.Regex{Pattern: metadata.ParentResourceId, Options: "i"}}
				andFilters = append(andFilters, bson.M{"metadata.parentResourceId": textFilter})
			}
		} else if key == "timestamp" && value[0] != "" {
			startTime := utils.GetFormattedStartTime(value[0])
			endTime := utils.GetFormattedEndTime(value[0])
			andFilters = append(andFilters, bson.M{
				"timestamp": bson.M{
					"$gte": startTime,
					"$lte": endTime,
				},
			})
		} else {
			if value[0] != "" {
				textFilter := bson.M{"$regex": primitive.Regex{Pattern: value[0], Options: "i"}}
				andFilters = append(andFilters, bson.M{key: textFilter})
			}
		}
	}

	// Use $and operator to combine multiple conditions
	if len(andFilters) > 0 {
		filter["$and"] = andFilters
	}
	return filter
}