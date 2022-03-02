package main

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	url      []string
	timer    int
	adresses string
)

func init() {
	adresses = os.Getenv("ADRESSES")
	delay := os.Getenv("DELAY")
	url = strings.Split(adresses, ",")

	frequency, err := strconv.Atoi(delay)

	if err != nil {
		log.Fatal("addresses should be set array as string")
	}

	if len(url) < 1 {
		log.Fatal("adresses lenght cannot less than 1")
	}

	if frequency < 0 {
		timer = 5
	} else {
		timer = frequency
	}
}

func main() {
	var ticker time.Duration
	ticker = time.Duration(timer)

	for range time.Tick(time.Second * ticker) {
		for _, v := range url {
			resp, err := http.Get(v + "/healthz")

			if err != nil {
				log.Fatal(err)
			}

			body := resp.Body

			data, err := io.ReadAll(body)

			if err != nil {
				log.Fatal(err)
			}

			log.Info(time.Now().Format("01-02-2006 15:04:05"), " ", string(data))
		}
	}
}
