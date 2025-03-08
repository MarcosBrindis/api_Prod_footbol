package application

import (
	"GameSport/src/game/domain/model"
	"GameSport/src/game/domain/ports"
	"context"
	"errors"
)

type UpdateGameUsecase struct {
	Repository ports.GameRepository
}

func (uc *UpdateGameUsecase) Execute(ctx context.Context, id int, game *model.Game) error {
	if game.Status != "live" && game.Status != "Finalized" {
		return errors.New("only status 'live' or 'Finalized' can be updated")
	}
	return uc.Repository.Update(ctx, id, game)
}
