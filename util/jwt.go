package util

import (
	"time"

	"github.com/ShawnRong/bento/logging"

	"github.com/dgrijalva/jwt-go"
)

//var jwtSecret = []byte(config.GetConfig().GetString("site.jwt"))
var jwtSecret = []byte("123321")

type Claims struct {
	Username string `json: "username"`
	Password string `json: "password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "bento",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	//integrate Log
	if err != nil {
		logging.Info(err.Error())
	}

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
