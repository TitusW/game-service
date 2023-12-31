package user

import (
	"context"
	"time"

	"github.com/TitusW/game-service/internal/entity"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

func convertCreateToModel(input entity.User) User {
	return User{
		Ksuid:      ksuid.New().String(),
		Email:      input.Email,
		Password:   input.Password,
		InsertedAt: time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func convertUpdateToModel(input entity.User) User {
	return User{
		Ksuid:     ksuid.New().String(),
		Email:     input.Email,
		Password:  input.Password,
		UpdatedAt: time.Now(),
	}
}

func convertToEntity(model User) entity.User {
	return entity.User{
		Ksuid:    model.Ksuid,
		Email:    model.Email,
		Password: model.Password,
	}
}

func convertToEntities(rows []User) []entity.User {
	var entities []entity.User
	for _, row := range rows {
		entities = append(entities, convertToEntity(row))
	}
	return entities
}

func (m UserModule) CreateTX(ctx context.Context, input entity.User, tx *gorm.DB) (entity.User, error) {
	user := convertCreateToModel(input)

	err := tx.WithContext(ctx).Create(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	returnUser := convertToEntity(user)

	return returnUser, nil
}

func (m UserModule) Update(ctx context.Context, input entity.User) (entity.User, error) {
	user := convertUpdateToModel(input)

	err := m.db.WithContext(ctx).Model(&user).Where("ksuid = ?", user.Ksuid).Updates(user).Error

	if err != nil {
		return entity.User{}, err
	}

	returnUser := convertToEntity(user)

	return returnUser, nil
}

func (m UserModule) Get(ctx context.Context, ksuid string) (entity.User, error) {
	var user User

	err := m.db.WithContext(ctx).Where("ksuid = ?", ksuid).First(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	returnUser := convertToEntity(user)

	return returnUser, nil
}

func (m UserModule) GetUsers(
	ctx context.Context,
	email string,
	bankAccountName string,
	bankAccountNumber string,
	bankName string,
	currentAmount float64,
	operator string,
) ([]entity.User, error) {
	var user []User

	err := m.db.WithContext(ctx).Scopes(
		filterByEmail(&email),
		joinBankAccount(),
		joinWallet(),
		filterByBankAccountName(&bankAccountName),
		filterByBankAccountNumber(&bankAccountNumber),
		filterByBankName(&bankName),
		filterByWalletCurrentAmount(&currentAmount, operator),
	).Find(&user).Error

	if err != nil {
		return []entity.User{}, err
	}

	returnUser := convertToEntities(user)

	return returnUser, nil
}

func (m UserModule) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var user User

	err := m.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	returnUser := convertToEntity(user)

	return returnUser, nil
}
