package application

import (
	"GameSport/src/sportEvent/domain/model"
	"GameSport/src/sportEvent/domain/ports"
	"context"
	"encoding/json"
)

type CreateSportsEventUsecase struct {
	Repository ports.SportsEventRepository
	Publisher  ports.EventPublisher
}

func NewCreateSportsEventUsecase(repo ports.SportsEventRepository, publisher ports.EventPublisher) *CreateSportsEventUsecase {
	return &CreateSportsEventUsecase{
		Repository: repo,
		Publisher:  publisher,
	}
}

func (uc *CreateSportsEventUsecase) Execute(ctx context.Context, event *model.SportsEvent) error {

	// Guardar en la base de datos
	if err := uc.Repository.Create(ctx, event); err != nil {
		return err
	}

	// Publicar en RabbitMQ
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return uc.Publisher.PublishEvent(ctx, string(eventJSON))
}
