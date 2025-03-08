package controller

import (
	"GameSport/src/sportEvent/application"
	"GameSport/src/sportEvent/domain/model"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateSportEventController struct {
	UseCase *application.CreateSportsEventUsecase
}

func NewCreateSportEventController(useCase *application.CreateSportsEventUsecase) *CreateSportEventController {
	return &CreateSportEventController{UseCase: useCase}
}

func (c *CreateSportEventController) HandleCreate(ctx *gin.Context) {
	var event model.SportsEvent
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	/*if err := c.UseCase.Execute(ctx.Request.Context(), &event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar o publicar el evento"})
		return
	}
	ctx.JSON(http.StatusCreated, event)*/

	reqContext := context.Background()
	if err := c.UseCase.Execute(reqContext, &event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Evento creado exitosamente"})
}
