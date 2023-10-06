package entity

type User struct {
	Ksuid    string `json:"ksuid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetail struct {
	Ksuid        string            `json:"ksuid"`
	Email        string            `json:"email"`
	BankAccounts []UserBankAccount `json:"bank_accounts"`
	Wallet       Wallet            `json:"wallet"`
}
