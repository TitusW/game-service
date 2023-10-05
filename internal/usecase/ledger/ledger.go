package ledger

import (
	"context"
	"fmt"

	"github.com/TitusW/game-service/internal/entity"
	"gorm.io/gorm"
)

func (uc Usecase) Topup(ctx context.Context, input entity.Ledger) (entity.Ledger, error) {
	var ledger entity.Ledger
	var wallet entity.Wallet
	fmt.Println("before_transaction")

	err := uc.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		fmt.Println("create_tx")
		createdLedger, err := uc.ledgerResource.CreateTX(ctx, input, tx)
		ledger = createdLedger
		if err != nil {
			return err
		}

		fmt.Println("get_tx")
		ledgers, err := uc.ledgerResource.GetTX(ctx, createdLedger.Ksuid, tx)
		wallet.CurrentAmount += calculateAmount(wallet, ledger)
		if err != nil {
			return err
		}

		for _, ledger := range ledgers {
			wallet.CurrentAmount += calculateAmount(wallet, ledger)
		}

		fmt.Println("update_tx")
		wallet.Ksuid = input.WalletKsuid
		_, err = uc.walletResource.UpdateTX(ctx, wallet, tx)
		if err != nil {
			return err
		}

		return nil
	})

	return ledger, err
}

func calculateAmount(wallet entity.Wallet, ledger entity.Ledger) float64 {
	switch {
	case ledger.Category == "MONEY_IN":
		wallet.CurrentAmount += ledger.Amount
	case ledger.Category == "MONEY_OUT":
		wallet.CurrentAmount -= ledger.Amount
	}

	return wallet.CurrentAmount
}
