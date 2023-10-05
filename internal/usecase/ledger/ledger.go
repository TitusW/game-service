package ledger

import (
	"context"

	"github.com/TitusW/game-service/internal/entity"
	"gorm.io/gorm"
)

func (uc Usecase) Topup(ctx context.Context, input entity.Ledger) (entity.Ledger, error) {
	var ledger entity.Ledger
	var wallet entity.Wallet

	err := uc.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		createdLedger, err := uc.ledgerResource.CreateTX(ctx, input, tx)
		ledger = createdLedger
		if err != nil {
			return err
		}

		ledgers, err := uc.ledgerResource.GetByWalletKsuidTX(ctx, createdLedger.WalletKsuid, tx)
		if err != nil {
			return err
		}

		for _, ledger := range ledgers {
			wallet.CurrentAmount = calculateAmount(wallet, ledger)
		}

		wallet.Ksuid = input.WalletKsuid
		_, err = uc.walletResource.UpdateTX(ctx, wallet, tx)
		if err != nil {
			return err
		}

		return nil
	})

	return ledger, err
}

func calculateAmount(currentAmount entity.Wallet, ledger entity.Ledger) float64 {
	switch {
	case ledger.Category == "MONEY_IN":
		currentAmount.CurrentAmount += ledger.Amount
	case ledger.Category == "MONEY_OUT":
		currentAmount.CurrentAmount -= ledger.Amount
	}

	return currentAmount.CurrentAmount
}
