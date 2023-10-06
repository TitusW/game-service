package user

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
	"gorm.io/gorm"
)

type UserModuleItf interface {
	CreateTX(ctx context.Context, input entity.User, tx *gorm.DB) (entity.User, error)
	Update(ctx context.Context, input entity.User) (entity.User, error)
	Get(ctx context.Context, ksuid string) (entity.User, error)
}

type WalletModuleItf interface {
	CreateTX(ctx context.Context, input entity.Wallet, tx *gorm.DB) (entity.Wallet, error)
	GetByUserKsuid(ctx context.Context, userKsuid string) (entity.Wallet, error)
}

type BankAccountModuleItf interface {
	GetByUserKsuid(ctx context.Context, userKsuid string) ([]entity.UserBankAccount, error)
}

type Usecase struct {
	userResource        UserModuleItf
	walletResource      WalletModuleItf
	bankAccountResource BankAccountModuleItf
	db                  *gorm.DB
}

func New(
	userResource UserModuleItf,
	walletResource WalletModuleItf,
	bankAccountResource BankAccountModuleItf,
	db *gorm.DB,
) Usecase {
	return Usecase{
		userResource:        userResource,
		walletResource:      walletResource,
		bankAccountResource: bankAccountResource,
		db:                  db,
	}
}
