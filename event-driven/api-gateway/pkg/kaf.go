package pkg

import (
	"crypto/tls"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	log "github.com/sirupsen/logrus"
	"os"
)

var KafkaDialer *kafka.Dialer

func init() {
	password := os.Getenv("KAFKA_PASS")
	username := os.Getenv("KAFKA_USERNAME")

	mechanism, err := scram.Mechanism(scram.SHA256, username, password)

	if err != nil {
		log.Fatal(err)
	}

	KafkaDialer = &kafka.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}
}
