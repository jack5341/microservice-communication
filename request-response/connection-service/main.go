package main

import (
	redis "connection-service/pkg"
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Person struct {
	Username string
	Follows  []string
}

// In-memory store
var DB = []Person{
	{
		Username: "anil",
		Follows: []string{
			"nedim",
		},
	},
	{
		Username: "nedim",
		Follows: []string{
			"anil",
		},
	},
	{
		Username: "kaan",
		Follows: []string{
			"nedim", "anil",
		},
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
	client := redis.Client()
	client.Set("language", "go", 0)
	log.Info("server is running on port", port)

	// Getting requests by endpoint "connections"
	http.HandleFunc("/connections", Connections)
	http.HandleFunc("/healthz", HealthCheck)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("connection service is alive"))
	if err != nil {
		return
	}
}

func Connections(w http.ResponseWriter, r *http.Request) {
	// Parsing username by query
	username := r.URL.Query().Get("username")

	// Looking for is exist this username in our DB
	for _, person := range DB {
		if username == person.Username {
			// If yes send his follows
			out, _ := json.Marshal(person.Follows)
			fmt.Fprintf(w, string(out))
		}
	}
}
