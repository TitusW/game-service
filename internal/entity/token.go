package entity

type Token struct {
	Ksuid   string
	Token   string
	Expired bool
	Revoked bool
}

type TokenResponse struct {
	Token        string
	RefreshToken string
}
