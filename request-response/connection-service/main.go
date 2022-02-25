package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Username string
	Follows  []string
}

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

func main() {
	http.HandleFunc("/connections", Connections)
	http.ListenAndServe(":8081", nil)
}

func Connections(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	for _, person := range DB {
		if username == person.Username {
			out, _ := json.Marshal(person.Follows)
			fmt.Fprintf(w, string(out))
		}
	}
}
