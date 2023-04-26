package pubsub

import "github.com/dimaglushkov/epam-xm-test-assignment/internal/core/ports"

type Kafka struct{}

func NewKafka() ports.PubSubPort {
	return Kafka{}
}
