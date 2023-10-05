package wallet

import (
	"context"
	"time"

	"github.com/TitusW/game-service/internal/entity"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

func convertCreateToModel(input entity.Wallet) Wallet {
	return Wallet{
		Ksuid:         ksuid.New().String(),
		UserKsuid:     input.UserKsuid,
		CurrentAmount: 0,
		InsertedAt:    time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func convertUpdateToModel(input entity.Wallet) Wallet {
	return Wallet{
		Ksuid:         input.Ksuid,
		CurrentAmount: input.CurrentAmount,
		UpdatedAt:     time.Now(),
	}
}

func convertToEntity(row Wallet) entity.Wallet {
	return entity.Wallet{
		Ksuid:         row.Ksuid,
		UserKsuid:     row.UserKsuid,
		CurrentAmount: row.CurrentAmount,
	}
}

func (m Module) UpdateTX(ctx context.Context, input entity.Wallet, tx *gorm.DB) (entity.Wallet, error) {
	wallet := convertUpdateToModel(input)

	err := tx.WithContext(ctx).Where("ksuid = ?", input.Ksuid).Updates(&wallet).Error
	if err != nil {
		return entity.Wallet{}, err
	}

	return convertToEntity(wallet), nil
}

func (m Module) CreateTX(ctx context.Context, input entity.Wallet, tx *gorm.DB) (entity.Wallet, error) {
	wallet := convertCreateToModel(input)

	err := tx.WithContext(ctx).Create(&wallet).Error
	if err != nil {
		return entity.Wallet{}, err
	}

	return convertToEntity(wallet), nil
}
