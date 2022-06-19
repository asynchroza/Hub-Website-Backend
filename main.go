package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// fmt, encoding/json, strconv
)

// Event Model
type Event struct {
	ID        int    `json:id`
	Title     string `json:"title"`
	Place     string `json: "place"`
	StartTime string `json: "startTime"`
	EndTime   string `json: "endTime"`
	ImageLink string `json: "imageLink"`
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Route Handlers / Endpoints
	// Events Endpoints
	router.HandleFunc("/api/events", getEvents).Methods("GET")             // get list of events
	router.HandleFunc("/api/events/{id}", getSpecificEvent).Methods("GET") // get specific event
	router.HandleFunc("/api/events/", createEvent).Methods("PUT")          // create a event

	// Articles Endpoints
	router.HandleFunc("/api/articles", getArticles).Methods("GET")             // get list of articles
	router.HandleFunc("/api/articles/{id}", getSpecificArticle).Methods("GET") // get specific article
	router.HandleFunc("/api/articles/", createArticle).Methods("PUT")          // create an article

	// Serve and log exceptions
	log.Fatal(http.ListenAndServe(":8000", router))

}
