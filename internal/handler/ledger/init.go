package ledger

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
)

type UsecaseItf interface {
	Topup(ctx context.Context, input entity.Ledger) (entity.Ledger, error)
}

type Handler struct {
	uc UsecaseItf
}

func New(uc UsecaseItf) Handler {
	return Handler{
		uc: uc,
	}
}
