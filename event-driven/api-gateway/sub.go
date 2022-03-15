package main

import (
	"api-gateway/pkg"
	"context"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func Subscribe(broker string, topic string) (kafka.Message, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		GroupID:  "",
		Topic:    topic,
		Dialer:   pkg.KafkaDialer,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer r.Close()

	log.Info("api-gateway start to consuming")

	message, err := r.ReadMessage(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	return message, err
}
