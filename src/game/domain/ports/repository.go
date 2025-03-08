package ports

import (
	"GameSport/src/game/domain/model"
	"context"
)

type GameRepository interface {
	Create(ctx context.Context, gameSport *model.Game) error
	GetByID(ctx context.Context, id int) (*model.Game, error)
	Update(ctx context.Context, id int, gameSport *model.Game) error
	GetAll(ctx context.Context) ([]*model.Game, error)
}
