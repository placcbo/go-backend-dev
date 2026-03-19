package main

import (
	"fmt"
	"log"
	"net/http"

	"project2-url-shortener/database"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("database: %v", err)
	}
	fmt.Println("✓ Connected to PostgreSQL")

	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", handleShorten)
	mux.HandleFunc("/stats/", handleStats)
	mux.HandleFunc("/", handleRedirect)

	fmt.Println("✓ Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}