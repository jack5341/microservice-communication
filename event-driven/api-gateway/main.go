package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	broker string
	topic  string
)

func init() {
	broker = os.Getenv("BROKER")
	topic = os.Getenv("TOPIC")
}

func main() {
	err := Publish(broker, topic, Event{Value: "hello value", Key: "hey key"})

	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := Subscribe("rocket-01.srvs.cloudkafka.com:9094", "vve4vunz-event")

		if err != nil {
			log.Fatal(err)
		}

		log.Info(string(msg.Value))
	}
}
