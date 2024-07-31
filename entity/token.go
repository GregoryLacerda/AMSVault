package entity

type Token struct {
	Token      string `json:"acces_token"`
	Expiration int    `json:"expiration"`
}
