package ports

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type CompanyMutationEvent struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Time     time.Time `json:"time"`
	Producer string    `json:"producer"`
	Data     any       `json:"data"`
}

func NewCompanyMutationEvent(eventName string, producer string, data any) *CompanyMutationEvent {
	return &CompanyMutationEvent{
		Id:       uuid.New(),
		Name:     eventName,
		Time:     time.Now(),
		Producer: producer,
		Data:     data,
	}
}

type EventsWriter interface {
	Write(ctx context.Context, data ...any) error
	Close()
}
