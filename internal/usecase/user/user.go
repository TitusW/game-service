package user

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (uc Usecase) Register(ctx context.Context, input entity.User) (entity.User, error) {
	var user entity.User

	err := uc.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		password, err := hashPassword(input.Password)
		if err != nil {
			return err
		}

		input.Password = password
		createdUser, err := uc.userResource.CreateTX(ctx, input, tx)
		user = createdUser
		if err != nil {
			return err
		}

		newWallet := entity.Wallet{
			UserKsuid: createdUser.Ksuid,
		}

		_, err = uc.walletResource.CreateTX(ctx, newWallet, tx)
		if err != nil {
			return err
		}

		return nil
	})

	return user, err
}

func (uc Usecase) Update(ctx context.Context, input entity.User) (entity.User, error) {
	user, err := uc.userResource.Update(ctx, input)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
