package ledger

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
	"gorm.io/gorm"
)

type LedgerModuleItf interface {
	CreateTX(ctx context.Context, input entity.Ledger, tx *gorm.DB) (entity.Ledger, error)
	GetTX(ctx context.Context, walletKsuid string, tx *gorm.DB) ([]entity.Ledger, error)
}

type WalletModuleItf interface {
	UpdateTX(ctx context.Context, input entity.Wallet, tx *gorm.DB) (entity.Wallet, error)
}

type Usecase struct {
	ledgerResource LedgerModuleItf
	walletResource WalletModuleItf
	db             *gorm.DB
}

func New(
	ledgerResource LedgerModuleItf,
	walletResource WalletModuleItf,
	db *gorm.DB,
) Usecase {
	return Usecase{
		ledgerResource: ledgerResource,
		walletResource: walletResource,
		db:             db,
	}
}
