package user

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
)

func (uc Usecase) Register(ctx context.Context, input entity.User) (entity.User, error) {
	user, err := uc.resource.Create(ctx, input)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (uc Usecase) Update(ctx context.Context, input entity.User) (entity.User, error) {
	user, err := uc.resource.Update(ctx, input)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
