package utils

import (
	"fmt"
	"time"
)

type Log struct {
	Level      string   `json:"level" bson:"level"`
	Message    string   `json:"message" bson:"message"`
	ResourceId string   `json:"resourceId" bson:"resourceId"`
	Timestamp  string   `json:"timestamp" bson:"timestamp"`
	TraceId    string   `json:"traceId" bson:"traceId"`
	SpanId     string   `json:"spanId" bson:"spanId"`
	Commit     string   `json:"commit" bson:"commit"`
	Metadata   Metadata `json:"metadata" bson:"metadata"`
}

type Metadata struct {
	ParentResourceId string `json:"parentResourceId" bson:"parentResourceId"`
}

var (
	Database   string = "DyteLogs"
	Collection string = "logsCollection"
	dateLayout string = "2006-01-02"
)

func GetFormattedStartTime(dateTime string) string {
	startDate, err := time.Parse(dateLayout, dateTime)
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
		return ""
	}
	startTime := startDate.UTC()
	formattedStartTime := startTime.Format(time.RFC3339)
	return formattedStartTime
}
func GetFormattedEndTime(dateTime string) string {
	endDate, err := time.Parse(dateLayout, dateTime)
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
		return ""
	}
	startTime := endDate.UTC()
	endTime := startTime.Add(24*time.Hour - 1*time.Second)
	formattedEndTime := endTime.Format(time.RFC3339)
	return formattedEndTime
}
