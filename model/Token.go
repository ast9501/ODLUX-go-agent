package model

type Token struct {
	Token	string	`json:"access_token"`
	ExpiredTime	int64	`json:"expires_at"`
	Type	string	`json:"token_type"`
}