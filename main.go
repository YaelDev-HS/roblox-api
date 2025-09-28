package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	start()

	// Handlers
	mux.HandleFunc("/", getUsers)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port " + port)
	log.Println("PORT from env:", os.Getenv("PORT"))

	http.ListenAndServe(":"+port, mux)
}
