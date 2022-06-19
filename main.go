package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// fmt, encoding/json, strconv
)

// Event Model
type Event struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Place     string `json:"place"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	ImageLink string `json:"imageLink"`
}

// Article Model
type Article struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Author     *Author  `json: "author"`
	Paragraphs []string `json: "paragraphs"`
}

// Author
type Author struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	LinkedIn  string `json:"linkedin"`
}

// Slice of Articles
var articles []Article

// Get Events
func getEvents(w http.ResponseWriter, r *http.Request) {

}

func getEvent(w http.ResponseWriter, r *http.Request) {

}

func createEvent(w http.ResponseWriter, r *http.Request) {

}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {

}

func createArticle(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Mock data - implement DB
	articles = append(articles, Article{ID: "1", Title: "Welcome to The Hub", Author: &Author{Firstname: "Michael", Lastname: "Bozhilov"}, Paragraphs: []string{
		"sdfhsdjkfsdjkhfh", "sdjafhjsdfhkjsdfhj", "sdjkfhsdjfhjks",
	}})

	// Init Router
	router := mux.NewRouter()

	// Route Handlers / Endpoints
	// Events Endpoints
	router.HandleFunc("/api/events", getEvents).Methods("GET")     // get list of events
	router.HandleFunc("/api/events/{id}", getEvent).Methods("GET") // get specific event
	router.HandleFunc("/api/events/", createEvent).Methods("POST") // create a event

	// Articles Endpoints
	router.HandleFunc("/api/articles", getArticles).Methods("GET")     // get list of articles
	router.HandleFunc("/api/articles/{id}", getArticle).Methods("GET") // get specific article
	router.HandleFunc("/api/articles/", createArticle).Methods("POST") // create an article

	// Serve and log exceptions
	log.Fatal(http.ListenAndServe(":8000", router))

}
