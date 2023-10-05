package entity

type UserBankAccount struct {
	Ksuid           string `json:"ksuid"`
	UserKsuid       string `json:"user_ksuid"`
	BankAccountName string `json:"bank_account_name"`
	BankName        string `json:"bank_name"`
	AccountNumber   string `json:"account_number"`
}
