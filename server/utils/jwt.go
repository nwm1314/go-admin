package utils

import (
	"github.com/dgrijalva/jwt-go"
	"go-admin/server/model/request"
	"log"
)

var key = []byte("7ee47d22-7387-4f3e-8db4-1744d8fd3423")

func GenerateToken(claims request.BaseClaims) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		log.Println(err.Error())
	}
	return token, err
}

func ParseToken(tokenString string) (*request.BaseClaims, error) {
	Claims := &request.BaseClaims{}
	tokenClaims, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return key, nil
	})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*request.BaseClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
