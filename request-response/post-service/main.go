package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func main() {
	http.HandleFunc("/posts", Posts)
	http.ListenAndServe(":8082", nil)
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
