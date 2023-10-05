package ledger

import (
	"net/http"

	"github.com/TitusW/game-service/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h Handler) Topup(ctx *gin.Context) {
	var ledger entity.Ledger
	if err := ctx.ShouldBindJSON(&ledger); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ledger, err := h.uc.Topup(ctx, ledger)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": ledger,
	})
	return
}
