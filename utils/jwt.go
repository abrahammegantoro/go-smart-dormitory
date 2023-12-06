package utils

import "github.com/dgrijalva/jwt-go/v4"

var SecretKey = "$2a$10$W37qJj4ivbcTNhpMCpBBL.AtBiTL28SR0euqyPSiG73PfJlTcvfsm"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return webToken, nil
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
