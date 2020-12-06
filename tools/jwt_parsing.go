package main

import (
	"fmt"
	"strings"

	"github.com/MSDO-ImageHost/Posts/internal/auth"
	"github.com/dgrijalva/jwt-go"
)

func main() {

	tokenString := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiI1Iiwicm9sZSI6ImFkbWluIiwiaXNzIjoiSW1hZ2VIb3N0LnNkdS5kayIsImV4cCI6MTYzODU1NDMzNSwiaWF0IjoxNjA3MDE4MzM1fQ.zjBYhUKcGvWtZ-eTwVkOe-7vB9Fz0sb_Iqin290mhzw"

	segments := strings.Split(tokenString, ".")

	claims, err := jwt.DecodeSegment(segments[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(claims))
	fmt.Println(auth.DecodeClaims(claims))
}
