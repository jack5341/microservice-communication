package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Username string
	Post     string
}

// In-memory store
var DB = []Person{
	{
		Username: "anil",
		Post:     "microservices are good",
	},
	{
		Username: "nedim",
		Post:     "monoliths are good",
	},
	{
		Username: "kaan",
		Post:     "nothing is good",
	},
}

var (
	port string
)

func init() {
	// Getting port number from environment
	port = os.Getenv("PORT")

	if port == "" {
		// Default port is 8081
		port = "8082"
	}
}

func main() {
	flag.Parse()
	port := fmt.Sprintf(":%s", port)

	fmt.Println("Server is running on port", port)

	// Getting requests by endpoint "/posts"
	http.HandleFunc("/posts", Posts)
	http.HandleFunc("/healthz", HealthCheck)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post service is alive"))
}

func Posts(w http.ResponseWriter, r *http.Request) {
	// Parsing username by query
	username := r.URL.Query().Get("username")

	// Looking for is exist this username in our DB
	for _, person := range DB {
		if username == person.Username {
			// If yes send his posts
			out, _ := json.Marshal(person.Post)
			fmt.Fprintf(w, string(out))
		}
	}
}
