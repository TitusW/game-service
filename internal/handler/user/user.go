package user

import (
	"net/http"
	"strconv"

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

func (h Handler) GetUsers(ctx *gin.Context) {
	var users []entity.User
	var email string = ctx.Query("email")
	var bankAccountName string = ctx.Query("bankAccountName")
	var bankAccountNumber string = ctx.Query("bankAccountNumber")
	var bankName string = ctx.Query("bankName")
	var currentAmountStr string = ctx.Query("currentAmount")
	var operator string = ctx.Query("amountOperator")

	currentAmount, err := strconv.ParseFloat(currentAmountStr, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	users, err = h.uc.GetUsers(ctx, email, bankAccountName, bankAccountNumber, bankName, currentAmount, operator)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
