package domain

import (
	"strconv"
	"vid-outbox-demo-go/internal"
)

func generateEvent(item *Game, eventType string) *internal.Event {
	payload := gameEventData{
		Id:          item.EntityID,
		Name:        item.Name,
		Summary:     item.Summary,
		Url:         item.Url,
		ReleaseDate: item.ReleaseDate,
	}
	return internal.New(
		strconv.Itoa(int(item.EntityID)),
		eventType,
		payload.aggregateType(),
		payload,
	)
}
