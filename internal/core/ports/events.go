package ports

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CompanyMutationEvent struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Time     time.Time `json:"time"`
	Producer string    `json:"producer"`
	Data     any       `json:"data"`
}

func NewCompanyMutationEvent(eventName string, producer string, data any) *CompanyMutationEvent {
	return &CompanyMutationEvent{
		ID:       uuid.New(),
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
