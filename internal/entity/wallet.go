package entity

type Wallet struct {
	Ksuid         string  `json:"ksuid"`
	UserKsuid     string  `json:"user_ksuid"`
	CurrentAmount float64 `json:"current_amount"`
}
