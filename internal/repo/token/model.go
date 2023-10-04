package token

type Token struct {
	Token   string
	Expired bool
	Revoked bool
}
