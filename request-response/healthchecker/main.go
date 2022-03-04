package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	url       []string
	timer     int
	addresses string
)

func init() {
	addresses = os.Getenv("ADRESSES")
	delay := os.Getenv("DELAY")
	url = strings.Split(addresses, ",")

	frequency, _ := strconv.Atoi(delay)

	if len(url) < 1 {
		log.Fatal("addresses lenght cannot less than 1")
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
			go Caller(v)
		}
	}
}

func Caller(host string) {
	host = "http://" + host + "/healthz"
	resp, err := http.Get(host)

	if err != nil {
		fmt.Println(host)
		log.Warn(time.Now().Format("01-02-2006 15:04:05"), " ", "an error occured!")
		return
	}

	body := resp.Body

	data, err := io.ReadAll(body)

	if resp.StatusCode != 200 {
		log.Warn(time.Now().Format("01-02-2006 15:04:05"), " ", "still waiting response...")
		return
	}

	log.Info(time.Now().Format("01-02-2006 15:04:05"), " ", string(data))

	err = resp.Body.Close()
	if err != nil {
		return
	}
}
