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

var (
	port string
)

func init() {
	port = os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}
}

func main() {
	flag.Parse()
	fmt.Println(port)
	port := fmt.Sprintf(":%s", port)

	if port == "" {
		port = "8081"
	}

	fmt.Println(port)

	http.HandleFunc("/connections", Connections)
	http.ListenAndServe(port, nil)
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
