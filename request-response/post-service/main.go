package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
		port = "8081"
	}
}

func main() {
	flag.Parse()
	port := fmt.Sprintf(":%s", port)

	// Getting requests by endpoint "posts"
	http.HandleFunc("/posts", Posts)
	http.ListenAndServe(port, nil)
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
