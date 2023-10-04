package entity

type User struct {
	Ksuid    string `json:"ksuid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
