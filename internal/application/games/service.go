package games

import (
	"vid-outbox-demo-go/internal"
	"vid-outbox-demo-go/internal/domain"
	"vid-outbox-demo-go/internal/persistence"
)

type service struct {
	persistence.Connection
}

type Service interface {
	Create(aggregate internal.Aggregate) (*domain.Game, error)
	GetAll() ([]*domain.Game, error)
}

func NewService(conn persistence.Connection) Service {
	return &service{conn}
}

func (s *service) Create(aggregate internal.Aggregate) (game *domain.Game, err error) {
	value, ok := aggregate.(*domain.Game)
	if !ok {
		return nil, internal.ErrInvalidAggregate
	}

	value.AddEvent(domain.GameCreatedEvent(value))

	s.Select("EntityID", "Name", "Summary", "Url", "ReleaseDate").
		Create(value)
	return value, nil
}

func (s *service) GetAll() ([]*domain.Game, error) {
	var games []*domain.Game
	s.Find(&games)
	return games, nil
}
