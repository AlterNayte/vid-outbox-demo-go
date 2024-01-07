package internal

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

type Aggregate interface {
	GetEvents() []*Event
	AddEvent(event *Event)
	RemoveEvent(event *Event)
	ClearEvents()
}

type Event struct {
	ID            uuid.UUID `db:"id"`
	AggregateId   string    `db:"aggregateid"`
	AggregateType string    `db:"aggregatetype"`
	EventType     string    `db:"type"`
	Payload       *[]byte   `db:"payload"`
}

func (Event) TableName() string {
	return "outbox"
}

func New(aggregateId string, eventType string, aggregateType string, data interface{}) *Event {
	var payload *[]byte
	if d, err := json.Marshal(data); err != nil {
		panic(err)
	} else {
		payload = &d
	}
	return &Event{
		ID:            uuid.New(),
		AggregateId:   aggregateId,
		AggregateType: aggregateType,
		EventType:     eventType,
		Payload:       payload,
	}
}

var ErrInvalidAggregate = fmt.Errorf("invalid aggregate")
