package events

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaWriter struct {
	appName string
	brokers []string
	topic   string
	writer  *kafka.Writer
}

func NewKafkaWriter(brokers []string, topic string) *KafkaWriter {
	_, err := kafka.DialLeader(context.Background(), "tcp", brokers[0], topic, 0)
	if err != nil {
		log.Fatal(err)
	}
	return &KafkaWriter{
		brokers: brokers,
		topic:   topic,
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (kw *KafkaWriter) Write(ctx context.Context, data ...any) error {
	messages := make([]kafka.Message, 0, len(data))

	for _, d := range data {
		serializedData, err := json.Marshal(d)
		if err != nil {
			return err
		}
		messages = append(messages, kafka.Message{Value: serializedData})
	}

	err := kw.writer.WriteMessages(ctx, messages...)
	return err
}

func (kw *KafkaWriter) Close() {
	kw.writer.Close()
}
