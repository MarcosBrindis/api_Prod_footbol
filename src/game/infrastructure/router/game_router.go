package router

import (
	"GameSport/src/game/infrastructure/http/controller"

	"github.com/gin-gonic/gin"
)

func SetupGameRoutes(r *gin.Engine,
	create *controller.CreateGameController,
	get *controller.GetGameController,
	update *controller.UpdateGameController,
	getAll *controller.GetAllGameController,
) {

	gameGroup := r.Group("/Game")
	{
		gameGroup.POST("/", create.HandleCreate)
		gameGroup.GET("/:id", get.HandleGet)
		gameGroup.PUT("/:id", update.HandleUpdate)
		gameGroup.GET("/", getAll.HandleGetAll)
	}
}
