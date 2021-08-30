package auth

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/realwebdev/Bilal/clockify3/conf"
)

func GenerateJWT(useremail string) (string, error) {
	Accesskey := conf.Conf{}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = useremail
	claims["expirey"] = time.Now().Add(time.Minute * 15).Unix()

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tok.SignedString([]byte(Accesskey.ACCESS))
	if err != nil {
		log.Printf("Somethign went wrong %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func RefreshJWT(email string) (string, error) {
	RefreshKey := conf.Conf{}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["expirey"] = time.Now().Add(time.Minute * 200).Unix()

	tokenString, err := token.SignedString([]byte(RefreshKey.ACCESS))
	if err != nil {
		log.Printf("Somethign went wrong %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
