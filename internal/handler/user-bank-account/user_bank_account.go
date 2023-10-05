package userbankaccount

import (
	"net/http"

	"github.com/TitusW/game-service/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h Handler) Register(ctx *gin.Context) {
	var userBankAccount entity.UserBankAccount

	if err := ctx.ShouldBindJSON(&userBankAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userBankAccount, err := h.uc.Register(ctx, userBankAccount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": userBankAccount,
	})
	return
}
