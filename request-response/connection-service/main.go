package main

import "C"
import (
	db "connection-service/pkg"
	"connection-service/pkg/models"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
)

type Person struct {
	Username string
	Follows  []string
}

var Persons []Person

var (
	port   string
	DB     *gorm.DB
	client *redis.Client
)

func init() {
	// Getting port number from environment
	port = os.Getenv("PORT")

	if port == "" {
		// Default port is 8081
		port = "8081"
	}

	DB = db.Conn()
}

func main() {
	flag.Parse()
	port := fmt.Sprintf(":%s", port)
	log.Info("server is running on port", port)

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

	err = json.Unmarshal(obj, &Persons)

	if err != nil {
		log.Error(err)
	}

	for _, v := range Persons {
		isExist := DB.First(&models.Person{}, "username = ?", v.Username)

		if isExist.Error == nil {
			continue
		}

		if result := DB.Create(&models.Person{
			Username: v.Username,
			Follows:  strings.Join(v.Follows, ", "),
		}); result.Error != nil {
			log.Fatal(result.Error)
		}
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

	defer r.Body.Close()

	if err != nil {
		log.Error(err)
	}
}

func Connections(w http.ResponseWriter, r *http.Request) {
	// Parsing username by query
	username := r.URL.Query().Get("username")

	defer r.Body.Close()

	var user Person

	result := DB.Find(&user, Person{Username: username})

	if result.Error == nil {
		log.Warning("username couldn't find at connection-service db")
		fmt.Fprint(w, "this username is invalid")
	}

	fmt.Println(&user)

	/*
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
	*/
}
