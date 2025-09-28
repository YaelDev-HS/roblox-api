package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	port := os.Getenv("PORT")

	mux.HandleFunc("/", getUsers)

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on http://localhost:" + port)

	http.ListenAndServe(":"+port, mux)
}
