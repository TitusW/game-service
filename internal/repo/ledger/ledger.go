package ledger

import (
	"context"
	"time"

	"github.com/TitusW/game-service/internal/entity"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

func convertCreateToModel(input entity.Ledger) Ledger {
	return Ledger{
		Ksuid:       ksuid.New().String(),
		WalletKsuid: input.WalletKsuid,
		Category:    input.Category,
		Amount:      input.Amount,
		InsertedAt:  time.Now(),
	}
}

func convertToEntity(row Ledger) entity.Ledger {
	return entity.Ledger{
		Ksuid:       row.Ksuid,
		WalletKsuid: row.WalletKsuid,
		Category:    row.Category,
		Amount:      row.Amount,
	}
}
func convertToEntities(rows []Ledger) []entity.Ledger {
	var entities []entity.Ledger
	for _, row := range rows {
		entities = append(entities, convertToEntity(row))
	}
	return entities
}

func (m Module) CreateTX(ctx context.Context, input entity.Ledger, tx *gorm.DB) (entity.Ledger, error) {
	ledger := convertCreateToModel(input)

	err := tx.WithContext(ctx).Create(&ledger).Error
	if err != nil {
		return entity.Ledger{}, err
	}

	return convertToEntity(ledger), nil
}

func (m Module) GetByWalletKsuidTX(ctx context.Context, walletKsuid string, tx *gorm.DB) ([]entity.Ledger, error) {
	var ledgers []Ledger

	err := tx.WithContext(ctx).Where("wallet_ksuid", walletKsuid).Find(&ledgers).Error
	if err != nil {
		return []entity.Ledger{}, err
	}

	return convertToEntities(ledgers), nil
}
