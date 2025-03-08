package main

import (
	gameInfra "GameSport/src/game/infrastructure"
	gameRouter "GameSport/src/game/infrastructure/router"
	sportEventInfra "GameSport/src/sportEvent/infrastructure"
	sportEventRouter "GameSport/src/sportEvent/infrastructure/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Inicializar dependencias
	gameInfra.InitDependencies()
	sportEventInfra.InitDependencies()

	// Crear instancia de Gin
	r := gin.Default()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true }, // Permite cualquier origen
		MaxAge:           12 * time.Hour,
	}))
	// Configurar rutas con los controladores globales
	gameRouter.SetupGameRoutes(r, gameInfra.CreateGameController, gameInfra.GetGameController, gameInfra.UpdateGameController, gameInfra.GetAllGameController)
	sportEventRouter.SetupSportEventRoutes(r, sportEventInfra.CreateSportEventController, sportEventInfra.GetSportEventController)

	// Iniciar el servidor
	r.Run(":8080")
}
