package main

import (
	"api-gateway/pkg"
	"context"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type Event struct {
	Key   string
	Value string
}

func Publish(broker string, topic string, event Event) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
		Dialer:  pkg.KafkaDialer,
	})

	defer func() { _ = w.Close() }()

	msg := kafka.Message{
		Key:   []byte(event.Key),
		Value: []byte(event.Value),
	}

	err := w.WriteMessages(context.Background(), msg)

	log.Info("event " + event.Key + " is sent")

	if err != nil {
		return err
	}

	return nil
}
