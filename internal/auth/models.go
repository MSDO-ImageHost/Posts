package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Token *jwt.Token

type Claims struct {
	// Standard fields
	Issuer    string `mapstructure:"iss"`
	Subject   string `mapstructure:"sub"` // user id
	Audience  string `mapstructure:"aud"`
	ExpiresAt int64  `mapstructure:"exp"`
	NotBefore int64  `mapstructure:"nbf"`
	IssuedAt  int64  `mapstructure:"iat"`
	Id        string `mapstructure:"jti"`

	// Custom fields
	Rank Rank `mapstructure:"role"`
}

type Rank byte

type User struct {
	JwtToken string
	UserID   string
	Rank     Rank
}
