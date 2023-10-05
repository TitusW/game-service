package entity

type Ledger struct {
	Ksuid       string  `json:"ksuid"`
	WalletKsuid string  `json:"wallet_ksuid"`
	Category    string  `json:"category"`
	Amount      float64 `json:"amount"`
}
