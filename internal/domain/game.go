package domain

import (
	"fmt"
	"time"
	"vid-outbox-demo-go/internal"
)

type Game struct {
	ID          int64             `json:"-" db:"id"`
	EntityID    int64             `json:"entity_id" db:"entity_id"`
	Name        string            `json:"name" db:"name"`
	Summary     string            `json:"summary" db:"summary"`
	Url         string            `json:"url" db:"url"`
	ReleaseDate time.Time         `json:"releaseDate" db:"release_date"`
	Events      []*internal.Event `json:"-" gorm:"-"`
}

func (g *Game) AddEvent(event *internal.Event) {
	g.Events = append(g.Events, event)
}

func (g *Game) RemoveEvent(event *internal.Event) {
	for i, v := range g.Events {
		if v.AggregateId == event.AggregateId && v.AggregateType == event.AggregateType {
			g.Events = append(g.Events[:i], g.Events[i+1:]...)
		}
	}
}

func (g *Game) GetEvents() []*internal.Event {
	return g.Events
}

func (g *Game) ClearEvents() {
	g.Events = nil
}

type gameEventData struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Summary     string    `json:"summary"`
	Url         string    `json:"url"`
	ReleaseDate time.Time `json:"releaseDate"`
}

func (e *gameEventData) version() string {
	return "v0.0.2"
}
func (e *gameEventData) aggregateType() string {
	return fmt.Sprintf("Game.%s", e.version())
}

func GameCreatedEvent(item *Game) *internal.Event {
	return generateEvent(item, "GameCreatedEvent")
}

func GameUpdatedEvent(item *Game) *internal.Event {
	return generateEvent(item, "GameUpdatedEvent")
}
