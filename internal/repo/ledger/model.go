package ledger

import (
	"time"

	"gorm.io/gorm"
)

type Ledger struct {
	Ksuid       string
	WalletKsuid string
	Category    string
	Amount      float64
	InsertedAt  time.Time
	UpdatedAt   time.Time
	DeletedAt   *gorm.DeletedAt
}
