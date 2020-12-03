package auth

//func ValidateToken(tokenString string) (bool, error) {
//
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		// Don't forget to validate the alg is what you expect:
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//
//		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
//		return []byte(os.Getenv("JWT_HMAC_SECRET")), nil
//	})
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		fmt.Println(claims["foo"], claims["nbf"])
//	} else {
//		fmt.Println(err)
//	}
//}

func Decompose(token string) (jwtc JWTSkeleton, err error) {

	jwtc = JWTSkeleton{
		UserID: "123-christian-id",
		Role:   "admin",
	}

	return jwtc, nil
}
