package controller

import (
	"GameSport/src/sportEvent/application"
	"context"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type GetSportEventController struct {
	UseCase *application.GetSportsEventUsecase
}

func NewGetSportEventController(useCase *application.GetSportsEventUsecase) *GetSportEventController {
	return &GetSportEventController{UseCase: useCase}
}

func (c *GetSportEventController) HandleGet(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número entero"})
		return
	}

	reqContext := context.Background()
	filmhub, err := c.UseCase.Execute(reqContext, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "evento no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, filmhub)
}
