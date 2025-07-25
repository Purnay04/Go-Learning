package utils

import "github.com/golang-jwt/jwt/v5"

var jwtKey = []byte("secretekey")

func CreateToken(claims jwt.MapClaims) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenStruct.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
