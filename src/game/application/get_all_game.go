package application

import (
	"GameSport/src/game/domain/model"
	"GameSport/src/game/domain/ports"
	"context"
)

type GetAllGameUsecase struct {
	Repository ports.GameRepository
}

func (uc *GetAllGameUsecase) Execute(ctx context.Context) ([]*model.Game, error) {
	games, err := uc.Repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return games, nil
}
