package userbankaccount

import "time"

type UserBankAccount struct {
	Ksuid           string
	UserKsuid       string
	BankAccountName string
	BankName        string
	AccountNumber   string
	InsertedAt      time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}
