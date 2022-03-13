package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	db "post-service/pkg"
)

type Person struct {
	Username string
	Post     string
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
		port = "8082"
	}
}

func main() {
	flag.Parse()
	port = fmt.Sprintf(":%s", port)

	fmt.Println("Server is running on port", port)

	// We're connect to Redis
	client = db.Conn()

	// Inserting inital datas
	obj, err := json.Marshal([]Person{
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
	})

	// Set inital values to Redis
	err = client.Set("id1234", obj, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// Getting requests by endpoint "/posts"
	http.HandleFunc("/posts", Posts)
	http.HandleFunc("/healthz", HealthCheck)

	err = http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(w http.
	ResponseWriter, r *http.Request) {
	w.Write([]byte("post service is alive"))
}

func Posts(w http.ResponseWriter, r *http.Request) {
	// Parsing username by query
	username := r.URL.Query().Get("username")

	// Getting our data by key
	val, err := client.Get("id1234").Result()
	if err != nil {
		log.Fatal(err)
	}

	// Created an []Person array cause of append
	var personList []Person

	// Unmarshal our string to json
	err = json.Unmarshal([]byte(val), &personList)

	if err != nil {
		log.Fatal(err)
	}

	// Looking for is exist this username in our DB
	for _, person := range personList {
		if username == person.Username {
			// If yes send his posts
			out, _ := json.Marshal(person.Post)
			fmt.Fprintf(w, string(out))
		} else {
			fmt.Fprintf(w, "invalid username")
		}
	}
}
