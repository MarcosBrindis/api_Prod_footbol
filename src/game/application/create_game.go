package application

import (
	"GameSport/src/game/domain/model"
	"GameSport/src/game/domain/ports"
	"context"
)

type CreateGameUsecase struct {
	Repository ports.GameRepository
}

func (uc *CreateGameUsecase) Execute(ctx context.Context, sportsEvent *model.Game) error {
	return uc.Repository.Create(ctx, sportsEvent)
}
