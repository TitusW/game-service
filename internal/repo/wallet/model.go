package wallet

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	Ksuid         string
	UserKsuid     string
	CurrentAmount float64
	InsertedAt    time.Time
	UpdatedAt     time.Time
	DeletedAt     *gorm.DeletedAt
}
