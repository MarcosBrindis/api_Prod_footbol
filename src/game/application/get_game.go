package application

import (
	"GameSport/src/game/domain/model"
	"GameSport/src/game/domain/ports"
	"context"
)

type GetGameUsecase struct {
	Repository ports.GameRepository
}

func (uc *GetGameUsecase) Execute(ctx context.Context, id int) (*model.Game, error) {
	return uc.Repository.GetByID(ctx, id)
}
