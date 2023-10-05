package userbankaccount

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
)

func (uc Usecase) Register(ctx context.Context, input entity.UserBankAccount) (entity.UserBankAccount, error) {
	userBankAccount, err := uc.resource.Register(ctx, input)
	if err != nil {
		return entity.UserBankAccount{}, err
	}

	return userBankAccount, nil
}
