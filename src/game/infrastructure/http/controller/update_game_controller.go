package controller

import (
	"GameSport/src/game/application"
	"GameSport/src/game/domain/model"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateGameController struct {
	UseCase *application.UpdateGameUsecase
}

func NewUpdateGameController(useCase *application.UpdateGameUsecase) *UpdateGameController {
	return &UpdateGameController{UseCase: useCase}
}

func (c *UpdateGameController) HandleUpdate(ctx *gin.Context) {
	var game model.Game
	if err := ctx.ShouldBindJSON(&game); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número entero"})
		return
	}

	reqContext := context.Background()
	if err := c.UseCase.Execute(reqContext, id, &game); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "juego actualizado exitosamente"})
}
