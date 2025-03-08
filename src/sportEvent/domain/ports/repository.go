package ports

import (
	"GameSport/src/sportEvent/domain/model"
	"context"
)

type SportsEventRepository interface {
	Create(ctx context.Context, sportsEvent *model.SportsEvent) error
	GetByID(ctx context.Context, id int) (*model.SportsEvent, error)
}
