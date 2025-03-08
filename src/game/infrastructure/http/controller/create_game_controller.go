package controller

import (
	"GameSport/src/game/application"
	"GameSport/src/game/domain/model"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateGameController struct {
	UseCase *application.CreateGameUsecase
}

func NewCreateGameController(useCase *application.CreateGameUsecase) *CreateGameController {
	return &CreateGameController{UseCase: useCase}
}

func (c *CreateGameController) HandleCreate(ctx *gin.Context) {
	var game model.Game
	if err := ctx.ShouldBindJSON(&game); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reqContext := context.Background()
	if err := c.UseCase.Execute(reqContext, &game); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "juego creado exitosamente"})
}
