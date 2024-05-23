package main

import (
	"log"
	"net/http"

	"github.com/Sneh16Shah/dytelogingestor/server/routeHandler"
)



func main() {
	http.HandleFunc("/", routeHandler.IngestLogHandler)
	http.HandleFunc("/queryLogs", routeHandler.QueryLogsHandler)
	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", addCorsHeaders(http.DefaultServeMux)))
}

func addCorsHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
