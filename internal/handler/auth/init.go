package auth

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
)

type UsecaseItf interface {
	Login(ctx context.Context, input entity.User) (entity.TokenResponse, error)
	Logout(ctx context.Context, tokenString string) error
}

type Handler struct {
	uc UsecaseItf
}

func New(uc UsecaseItf) Handler {
	return Handler{
		uc: uc,
	}
}
