package userbankaccount

import (
	"context"
	"time"

	"github.com/TitusW/game-service/internal/entity"
	"github.com/segmentio/ksuid"
)

func convertCreateToModel(input entity.UserBankAccount) UserBankAccount {
	return UserBankAccount{
		Ksuid:           ksuid.New().String(),
		UserKsuid:       input.UserKsuid,
		BankAccountName: input.BankAccountName,
		BankName:        input.BankName,
		AccountNumber:   input.AccountNumber,
		InsertedAt:      time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func convertToEntity(row UserBankAccount) entity.UserBankAccount {
	return entity.UserBankAccount{
		Ksuid:           row.Ksuid,
		UserKsuid:       row.UserKsuid,
		BankAccountName: row.BankAccountName,
		BankName:        row.BankName,
		AccountNumber:   row.AccountNumber,
	}
}

func convertToEntities(rows []UserBankAccount) []entity.UserBankAccount {
	var entities []entity.UserBankAccount
	for _, row := range rows {
		entities = append(entities, convertToEntity(row))
	}
	return entities
}

func (m Module) Register(ctx context.Context, input entity.UserBankAccount) (entity.UserBankAccount, error) {
	bankAccount := convertCreateToModel(input)

	err := m.db.WithContext(ctx).Create(&bankAccount).Error
	if err != nil {
		return entity.UserBankAccount{}, err
	}

	returnBankAccount := convertToEntity(bankAccount)

	return returnBankAccount, nil
}

func (m Module) GetByUserKsuid(ctx context.Context, userKsuid string) ([]entity.UserBankAccount, error) {
	var userBankAccounts []UserBankAccount

	err := m.db.WithContext(ctx).Where("user_ksuid = ?", userKsuid).Find(&userBankAccounts).Error
	if err != nil {
		return []entity.UserBankAccount{}, err
	}

	return convertToEntities(userBankAccounts), nil
}
