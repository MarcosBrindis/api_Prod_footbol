package application

import (
	"GameSport/src/sportEvent/domain/model"
	"GameSport/src/sportEvent/domain/ports"
	"context"
)

type GetSportsEventUsecase struct {
	Repository ports.SportsEventRepository
}

func (uc *GetSportsEventUsecase) Execute(ctx context.Context, id int) (*model.SportsEvent, error) {
	return uc.Repository.GetByID(ctx, id)
}
