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
	port = os.Getenv("PORT")

	if port == "" {
		port = "8082"
	}
}

func main() {
	flag.Parse()
	fmt.Println(port)
	port := fmt.Sprintf(":%s", port)

	if port == " " {
		log.Panic("port is required arg")
	}

	http.HandleFunc("/posts", Posts)
	http.ListenAndServe(port, nil)
}

func Posts(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	for _, person := range DB {
		if username == person.Username {
			out, _ := json.Marshal(person.Post)
			fmt.Fprintf(w, string(out))
		}
	}
}
