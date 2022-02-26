package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Person struct {
	Username string
	Follows  []string
	Post     string
}

var DB = []Person{
	{
		Username: "anil",
		Follows: []string{
			"nedim",
		},
		Post: "JS is good",
	},
	{
		Username: "nedim",
		Follows: []string{
			"anil",
		},
		Post: "Go is fast",
	},
	{
		Username: "kaan",
		Follows: []string{
			"nedim", "anil",
		},
		Post: "nothing is good",
	},
}

var (
	port string
)

func init() {
	port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
}

func main() {
	flag.Parse()
	fmt.Println(port)
	port := fmt.Sprintf(":%s", port)

	if port == "" {
		port = "8080"
	}

	fmt.Println(port)

	http.HandleFunc("/timeline", Timeline)
	http.HandleFunc("/connections", Connections)
	http.HandleFunc("/post", Post)
	http.ListenAndServe(port, nil)
}

func Timeline(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	timeline := make(map[string]string)

	for _, person := range DB {
		if username == person.Username {
			for _, followed := range person.Follows {
				timeline[followed] = getPostOfUser(followed)
			}
		}
	}

	out, _ := json.Marshal(timeline)
	fmt.Fprintf(w, string(out))
}

func Connections(w http.ResponseWriter, r *http.Request) {
	type Flw struct {
		Follower string `json:"follower"`
		Followed string `json:"followed"`
	}

	var flw Flw

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &flw)

	for i, person := range DB {
		if flw.Follower == person.Username {
			DB[i].Follows = append(DB[i].Follows, flw.Followed)
		}
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	type Post struct {
		Username string `json:"username"`
		Post     string `json:"post"`
	}

	var post Post

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal([]byte(string(body)), &post)

	if err != nil {
		fmt.Println(err)
	}

	for i, person := range DB {
		if post.Username == person.Username {
			DB[i].Post = post.Post
		}
	}
}

func getPostOfUser(username string) string {
	for _, person := range DB {
		if username == person.Username {
			return person.Post
		}
	}

	return ""
}
