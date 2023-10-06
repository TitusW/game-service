package user

import (
	"net/http"

	"github.com/TitusW/game-service/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h Handler) Register(ctx *gin.Context) {
	var user entity.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.uc.Register(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
	return
}

func (h Handler) Update(ctx *gin.Context) {
	var user entity.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.uc.Update(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	return
}

func (h Handler) GetUserDetails(ctx *gin.Context) {
	var userDetails entity.UserDetail
	var ksuid string = ctx.Param("ksuid")

	userDetails, err := h.uc.GetUserDetails(ctx, ksuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": userDetails,
	})
}
