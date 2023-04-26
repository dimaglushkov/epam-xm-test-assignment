package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"time"
)

func NewReader(brokers []string, topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topic,
		Partition: 0,
	})
}

func main() {
	time.Sleep(5 * time.Second)
	reader := NewReader([]string{os.Getenv("KAFKA_BROKERS")}, os.Getenv("KAFKA_TOPIC"))
	defer reader.Close()
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("event: %s", m.Value)
	}
}
