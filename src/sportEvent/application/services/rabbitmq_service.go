package services

import (
	"context"
	"encoding/json"

	"GameSport/src/sportEvent/application"
	"GameSport/src/sportEvent/domain/model"
)

// RabbitMQService encapsula el caso de uso para publicar eventos deportivos.
type RabbitMQService struct {
	publisherUsecase *application.CreateSportsEventUsecase
}

// NewRabbitMQService crea una nueva instancia de RabbitMQService inyectándole el caso de uso.
func NewRabbitMQService(publisherUsecase *application.CreateSportsEventUsecase) *RabbitMQService {
	return &RabbitMQService{publisherUsecase: publisherUsecase}
}

// SendMessage envía un mensaje utilizando el caso de uso.
func (svc *RabbitMQService) SendMessage(ctx context.Context, message string) error {
	var sportsEvent model.SportsEvent
	if err := json.Unmarshal([]byte(message), &sportsEvent); err != nil {
		return err
	}
	return svc.publisherUsecase.Execute(ctx, &sportsEvent)
}
