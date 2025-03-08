package controller

import (
	"GameSport/src/game/application"
	"context"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type GetGameController struct {
	UseCase *application.GetGameUsecase
}

func NewGetGameController(useCase *application.GetGameUsecase) *GetGameController {
	return &GetGameController{UseCase: useCase}
}

func (c *GetGameController) HandleGet(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número entero"})
		return
	}

	reqContext := context.Background()
	game, err := c.UseCase.Execute(reqContext, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "juego no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, game)
}
