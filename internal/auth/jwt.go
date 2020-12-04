package auth

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

func Parse(tokenString string) (Token, error) {
	var err error

	token, err := jwt.Parse(tokenString, FetchJwtSecret)
	if err != nil {
		return nil, fmt.Errorf("Invalid JWT")
	}

	// Token is valid
	if token.Valid {
		return (Token)(token), nil
	}
	return nil, err
}

func FetchJwtSecret(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_HMAC_SECRET")), nil
}

func DecodeClaims(claims map[string]interface{}) (result Claims, err error) {
	err = mapstructure.Decode(claims, &result)
	return result, err
}

/*
// Check for required header fields
tokenString, tokenPresent := msg.Headers["jwt"].(string)
if !tokenPresent {
	headers["status_code"] = http.StatusUnauthorized
	headers["status_code_msg"] = http.StatusText(http.StatusUnauthorized)
	log.Println(_LOG_TAG, "Rejected request with correlation id", msg.CorrelationId)
	if err := PublicateResponse(handleConf, msg, headers, nil, false, start); err != nil {
		log.Fatal(_LOG_TAG, "Failed process response to", msg.CorrelationId, err)
	}
	continue
}

// Parse JWT
token, err := auth.Parse(tokenString)
if err != nil || token == nil {
	headers["status_code"] = http.StatusUnauthorized
	headers["status_code_msg"] = http.StatusText(http.StatusUnauthorized)
	log.Println(_LOG_TAG, "Rejected request with correlation id", msg.CorrelationId)
	if err := PublicateResponse(handleConf, msg, headers, nil, false, start); err != nil {
		log.Fatal(_LOG_TAG, "Failed process response to", msg.CorrelationId, err)
	}
	continue
}*/
