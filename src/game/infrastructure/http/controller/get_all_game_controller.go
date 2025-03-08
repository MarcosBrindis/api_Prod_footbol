package controller

import (
	"GameSport/src/game/application"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllGameController struct {
	UseCase *application.GetAllGameUsecase
}

func NewGetAllGameController(useCase *application.GetAllGameUsecase) *GetAllGameController {
	return &GetAllGameController{UseCase: useCase}
}

func (c *GetAllGameController) HandleGetAll(ctx *gin.Context) {
	reqContext := context.Background()
	users, err := c.UseCase.Execute(reqContext)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
