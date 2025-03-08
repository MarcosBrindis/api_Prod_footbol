package infrastructure

import (
	db "GameSport/src/core/database/postgres"
	"GameSport/src/sportEvent/application"
	repo "GameSport/src/sportEvent/infrastructure/database/postgres"
	"GameSport/src/sportEvent/infrastructure/http/controller"
	"GameSport/src/sportEvent/infrastructure/messaging"
	"fmt"
)

var (
	CreateSportEventController *controller.CreateSportEventController
	GetSportEventController    *controller.GetSportEventController
)

func InitDependencies() {
	// Inicializar la base de datos
	database, err := db.NewDatabase()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Inicializar el repositorio de SportEvent
	sportEventRepo := repo.NewSportEventRepository(database)

	// RabbitMQ
	rabbitMQAdapter, err := messaging.NewRabbitMQ("amqp://MarcosDaniel:123456789a@54.146.109.24:5672/", "sport_events")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
	}

	// Casos de uso CORREGIDOS
	createSportEventUsecase := application.NewCreateSportsEventUsecase(sportEventRepo, rabbitMQAdapter)
	getSportEventUsecase := application.GetSportsEventUsecase{Repository: sportEventRepo}

	// Crear instancias de los controladores
	CreateSportEventController = controller.NewCreateSportEventController(createSportEventUsecase)
	GetSportEventController = controller.NewGetSportEventController(&getSportEventUsecase)
}
