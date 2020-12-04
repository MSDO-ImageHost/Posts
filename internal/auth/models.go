package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Token *jwt.Token

type Claims struct {
	Issuer     string `mapstructure:"iss"`
	Subject    string `mapstructure:"sub"`
	Audience   string `mapstructure:"aud"`
	Expiration int    `mapstructure:"exp"`
	NotBefore  int    `mapstructure:"nbf"`
	IssuedAt   string `mapstructure:"iat"`
	JWTID      string `mapstructure:"jti"`
	Role       string `mapstructure:"role"`
}
