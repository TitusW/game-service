package user

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
)

type UsecaseItf interface {
	Register(context.Context, entity.User) (entity.User, error)
	Update(context.Context, entity.User) (entity.User, error)
	GetUserDetails(ctx context.Context, ksuid string) (entity.UserDetail, error)
	GetUsers(
		ctx context.Context,
		email string,
		bankAccountName string,
		bankAccountNumber string,
		bankName string,
		currentAmount float64,
		operator string) ([]entity.User, error)
}

type Handler struct {
	uc UsecaseItf
}

func New(uc UsecaseItf) Handler {
	return Handler{
		uc: uc,
	}
}
