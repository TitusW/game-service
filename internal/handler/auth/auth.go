package auth

import (
	"net/http"
	"strings"

	"github.com/TitusW/game-service/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h Handler) Login(ctx *gin.Context) {
	var user entity.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	tokenResponse, err := h.uc.Login(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": tokenResponse,
	})
	return
}

func (h Handler) Logout(ctx *gin.Context) {
	var token string
	token = ctx.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", -1)

	err := h.uc.Logout(ctx, token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Logout successful",
	})
	return
}
