package user

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
)

type ModuleItf interface {
	Create(ctx context.Context, input entity.User) (entity.User, error)
	Update(ctx context.Context, input entity.User) (entity.User, error)
}

type Usecase struct {
	resource ModuleItf
}

func New(resource ModuleItf) Usecase {
	return Usecase{
		resource: resource,
	}
}
