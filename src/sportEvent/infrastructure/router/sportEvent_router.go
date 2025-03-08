// router/sport_event_routes.go
package router

import (
	"GameSport/src/sportEvent/infrastructure/http/controller"

	"github.com/gin-gonic/gin"
)

func SetupSportEventRoutes(
	r *gin.Engine,
	create *controller.CreateSportEventController,
	get *controller.GetSportEventController,

) {
	sportEventGroup := r.Group("/sportEvent")
	{
		sportEventGroup.POST("/", create.HandleCreate)
		sportEventGroup.GET("/:id", get.HandleGet)
	}

}
