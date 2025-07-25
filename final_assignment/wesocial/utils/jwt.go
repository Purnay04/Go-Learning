package utils

import "github.com/golang-jwt/jwt/v5"

var JwtKey = []byte("secretekey")

func CreateToken(claims jwt.MapClaims) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenStruct.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
