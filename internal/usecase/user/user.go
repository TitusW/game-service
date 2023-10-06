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

func (uc Usecase) GetUserDetails(ctx context.Context, ksuid string) (entity.UserDetail, error) {
	var userDetail entity.UserDetail

	user, err := uc.userResource.Get(ctx, ksuid)
	if err != nil {
		return userDetail, err
	}

	userDetail.Ksuid = user.Ksuid
	userDetail.Email = user.Email

	bankAccounts, err := uc.bankAccountResource.GetByUserKsuid(ctx, ksuid)
	if err != nil {
		return userDetail, err
	}

	for _, bankAccount := range bankAccounts {
		userDetail.BankAccounts = append(userDetail.BankAccounts, bankAccount)
	}

	wallet, err := uc.walletResource.GetByUserKsuid(ctx, ksuid)
	if err != nil {
		return userDetail, err
	}
	userDetail.Wallet.Ksuid = wallet.Ksuid
	userDetail.Wallet.UserKsuid = wallet.UserKsuid
	userDetail.Wallet.CurrentAmount = wallet.CurrentAmount

	return userDetail, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
