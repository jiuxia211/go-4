package utils

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTsecret = []byte("BTBT")

type Claims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}
type EmailClaims struct {
	Password string `json:"password"`
	Email    string `json:"emali"`
	jwt.StandardClaims
}

func GenerateToken(username string, uid uint, password string) (string, error) {
	nowtTime := time.Now()
	expireTime := nowtTime.Add(24 * time.Hour)
	myClaims := Claims{
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  nowtTime.Unix(),
			ExpiresAt: expireTime.Unix(),
			Id:        strconv.Itoa(int(uid)),
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	token, err := tokenClaim.SignedString(JWTsecret)
	return token, err
}
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if myClaims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return myClaims, nil
		}
	}
	return nil, err

}
func GenerateEmailToken(uid, email, password string) (string, error) {
	nowtTime := time.Now()
	expireTime := nowtTime.Add(10 * time.Minute)
	myClaims := EmailClaims{
		Email:    email,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  nowtTime.Unix(),
			ExpiresAt: expireTime.Unix(),
			Id:        uid,
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	token, err := tokenClaim.SignedString(JWTsecret)
	return token, err
}
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if myClaims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return myClaims, nil
		}
	}
	return nil, err

}
