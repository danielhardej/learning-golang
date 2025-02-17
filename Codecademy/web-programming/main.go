package main

import (
	"log"
	"net/http"
)

func main() {

	// Initialize the server and set up the routes
	mux := http.NewServeMux()

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("GET /menu", handleMenu)
	mux.HandleFunc("GET /reviews", handleReviews)
	mux.HandleFunc("GET /review", handleReviewForm)
	mux.HandleFunc("POST /submitReview", handleReviewSubmission)

	log.Println("Starting the server at http://localhost:4001")

	// Start the server
	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatal("Error occurred while starting the server:", err)
	}
}
