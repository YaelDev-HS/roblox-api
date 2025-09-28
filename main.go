package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()

	start()
	go executeForMinute()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port " + port)
	log.Println("PORT from env:", os.Getenv("PORT"))

	http.ListenAndServe(":"+port, mux)
}

func executeForMinute() {
	for {
		getUsers()
		time.Sleep(time.Minute * 1)
	}
}
