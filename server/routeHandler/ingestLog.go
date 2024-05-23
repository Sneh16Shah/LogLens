package routeHandler

import (
	"encoding/json"
	"net/http"
	"github.com/Sneh16Shah/dytelogingestor/server/utils"
	"github.com/Sneh16Shah/dytelogingestor/server/databasehandler"
)

func IngestLogHandler(w http.ResponseWriter, r *http.Request) {
	var logData utils.Log
	err := json.NewDecoder(r.Body).Decode(&logData) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	databasehandler.SaveToMongoDB(logData) 	// Save the log to MongoDB

	w.WriteHeader(http.StatusOK)
}
