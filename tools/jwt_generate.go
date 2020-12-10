package main

import (
	"fmt"

	"github.com/MSDO-ImageHost/Posts/internal/auth"
	"github.com/MSDO-ImageHost/Posts/internal/utils"
	"github.com/dgrijalva/jwt-go"
)

func main() {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "12",
		"role": 0,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("secret"))
	fmt.Println(tokenString, err)

	parsedToken, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{"sub": nil, "role": -1}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	fmt.Println(parsedToken)

	a, err := auth.AuthJWT(tokenString)
	fmt.Println(utils.PrettyFormatMap(a), err)

}
