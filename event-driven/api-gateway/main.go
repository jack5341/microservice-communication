package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

var (
	broker string
	title  string
	topic  string
	port   string
)

type ServiceRequest struct {
	User    string
	Newuser string
	Date    string
}

func init() {
	broker = os.Getenv("BROKER")
	topic = os.Getenv("TOPIC")
	title = os.Getenv("EVENT_TITLE")
	port = os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
}

func main() {
	fmt.Println("Listening on port", port)

	http.HandleFunc("/", handleRequest)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	newUser := r.URL.Query().Get("new-user")

	if user == "" {
		_, err := fmt.Fprintf(w, "user is required arg")

		if err != nil {
			log.Error(err)
		}

		return
	}

	if newUser == "" {
		_, err := fmt.Fprintf(w, "new-user is required arg")

		if err != nil {
			log.Error(err)
		}

		return
	}

	obj := ServiceRequest{
		User:    user,
		Newuser: newUser,
		Date:    time.Now().String(),
	}

	out, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	err = Publish(broker, topic, Event{Value: string(out), Key: title})

	if err != nil {
		log.Error(err)
		return
	}

	log.Info("new event committed from api-gateway")
}
