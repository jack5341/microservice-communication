package main

import (
	db "connection-service/pkg"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Person struct {
	Username string
	Follows  []string
}

var (
	port   string
	client *redis.Client
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
	log.Info("server is running on port", port)

	client = db.Conn()

	obj, err := json.Marshal([]Person{
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
	})

	err = client.Set("id1234", obj, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// Getting requests by endpoint "connections"
	http.HandleFunc("/connections", Connections)

	// Endpoint for health-check
	http.HandleFunc("/healthz", HealthCheck)

	err = http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("connection service is alive"))
	if err != nil {
		log.Fatal(err)
	}
}

func Connections(w http.ResponseWriter, r *http.Request) {
	// Parsing username by query
	username := r.URL.Query().Get("username")

	val, err := client.Get("id1234").Result()
	if err != nil {
		log.Fatal(err)
	}

	var personList []Person
	err = json.Unmarshal([]byte(val), &personList)

	if err != nil {
		log.Fatal(err)
	}

	// Looking for is exist this username in our DB
	for _, person := range personList {
		if username == person.Username {
			// If yes send his follows
			out, _ := json.Marshal(person.Follows)
			fmt.Fprintf(w, string(out))
		} else {
			fmt.Fprintf(w, "invalid username")
		}
	}
}
