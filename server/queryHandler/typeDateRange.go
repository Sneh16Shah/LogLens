package queryHandler

import (
	"net/url"

	"github.com/Sneh16Shah/dytelogingestor/server/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTimeRangeFilter(queryParams url.Values) bson.M {
	startTime := utils.GetFormattedStartTime(queryParams.Get("startDate"))
	endTime := utils.GetFormattedEndTime(queryParams.Get("endDate"))
	filter := bson.M{
		"timestamp": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
	}
	return filter
}
