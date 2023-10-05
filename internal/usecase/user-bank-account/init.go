package userbankaccount

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
)

type ModuleItf interface {
	Register(ctx context.Context, input entity.UserBankAccount) (entity.UserBankAccount, error)
}

type Usecase struct {
	resource ModuleItf
}

func New(resource ModuleItf) Usecase {
	return Usecase{
		resource: resource,
	}
}
