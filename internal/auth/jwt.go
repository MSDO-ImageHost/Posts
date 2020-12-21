package auth

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

// Validates a JWT token string and returns a user object with extracted claims
func AuthJWT(tokenString string) (auth User, err error) {

	auth.JwtToken = tokenString
	token, err := jwt.ParseWithClaims(tokenString, claimsModel, FetchJwtSecret)
	if err != nil {
		return auth, err
	}

	// Token is valid
	if token.Valid {
		claims, err := DecodeClaims(token.Claims)
		if err != nil {
			return auth, err
		}
		auth.UserID = claims.Subject
		auth.Rank = claims.Rank
		return auth, nil
	}
	return auth, nil
}

// Fetches the secret (that was used to generate the JWTs) used to validate JWT tokens
func FetchJwtSecret(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(os.Getenv("JWT_HMAC_SECRET")), nil
}

// Convert the claim map into a structure
func DecodeClaims(claims interface{}) (result Claims, err error) {
	err = mapstructure.Decode(claims, &result)
	return result, err
}
