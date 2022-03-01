package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
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
	port         string
	brokerString string
	credUser     string
	credPass     string
)

func init() {
	port = os.Getenv("PORT")
	brokerString = os.Getenv("BROKER_STRING")
	credUser = os.Getenv("CREDENTIAL_USERNAME")
	credPass = os.Getenv("CREDENTIAL_PASSWORD")

	if credUser == "" {
		log.Fatal("CREDENTIAL_USERNAME is not set")
	}

	if credPass == "" {
		log.Fatal("CREDENTIAL_PASSWORD is not set")
	}

	if brokerString == "" {
		log.Fatal("BROKER_STRING is not set")
	}

	if port == "" {
		log.Fatal("PORT must be set")
	}
}

func main() {
	flag.Parse()

	mechanism, err := scram.Mechanism(scram.SHA256, credUser, credPass)
	if err != nil {
		log.Fatalln(err)
	}

	dialer := &kafka.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerString},
		Topic:   "$TOPIC",
		Dialer:  dialer,
	})

	w.Close()

	http.HandleFunc("/timeline", Timeline)
	http.HandleFunc("/connections", Connections)
	http.HandleFunc("/post", Post)
	fmt.Println("Listening on port", port)
	err = http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
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
