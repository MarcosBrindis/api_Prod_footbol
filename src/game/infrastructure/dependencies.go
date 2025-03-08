package infrastructure

import (
	"fmt"

	"GameSport/src/game/application"
	repo "GameSport/src/game/infrastructure/database/postgres"
	"GameSport/src/game/infrastructure/http/controller"

	db "GameSport/src/core/database/postgres"
)

var (
	CreateGameController *controller.CreateGameController
	GetGameController    *controller.GetGameController
	UpdateGameController *controller.UpdateGameController
	GetAllGameController *controller.GetAllGameController
)

func InitDependencies() {
	// Inicializar la base de datos

	database, err := db.NewDatabase()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Inicializar el repositorio de FilmHub
	gameRepo := repo.NewGameRepository(database)

	// Crear los casos de uso
	createGameUsecase := application.CreateGameUsecase{Repository: gameRepo}
	getGameUsecase := application.GetGameUsecase{Repository: gameRepo}
	updateGameUsecase := application.UpdateGameUsecase{Repository: gameRepo}
	getAllGameUsecase := application.GetAllGameUsecase{Repository: gameRepo}

	// Crear instancias de los controladores
	CreateGameController = controller.NewCreateGameController(&createGameUsecase)
	GetGameController = controller.NewGetGameController(&getGameUsecase)
	UpdateGameController = controller.NewUpdateGameController(&updateGameUsecase)
	GetAllGameController = controller.NewGetAllGameController(&getAllGameUsecase)

}
