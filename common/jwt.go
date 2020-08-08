package common

import (
	"ginEsseential/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User_info) (string,error) {
	expirationTime :=time.Now().Add(7*24*time.Hour)

	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "govzzp.cn",
			Subject: "user token",
		},

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString , err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString,nil
}
func ParseToken(tokenString string) (*jwt.Token,*Claims,error) {
	claims := &Claims{}

	token,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) ( i interface{},  err  error) {
		return jwtKey,nil
	})
	return token,claims,err
}

