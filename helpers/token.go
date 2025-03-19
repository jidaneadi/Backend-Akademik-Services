package helpers

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(key []byte, claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyTokens(key []byte, tokensString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokensString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeTokens(tokenString string, key []byte) (jwt.MapClaims, error) {
	token, err := VerifyTokens(key, tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid tokens")
}
