package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/realwebdev/Bilal/clockify3/conf"
)

func AuthenticateToken(r *http.Request) error {
	token, err := VerifyToken(r)

	if err != nil {

		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {

		return err
	}

	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	conf := conf.Conf{}
	tokenString, err := ExtractToken(r)

	if err != nil {

		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unauthorize signin: %v", token.Header["alg"])
		}

		return []byte(conf.ACCESS), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractToken(r *http.Request) (string, error) {
	bearToken := r.Header.Get("Authorization")
	StrArr := strings.Split(bearToken, " ") //slice substring

	if len(StrArr) != 2 {

		return "", errors.New("Unable to fetch token from the request")
	}

	return StrArr[1], nil
}
