package lib

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Id int64 `json:"id,omitempty"`
	Uin int64 `json:"uin,omitempty"`
	Name string `json:"name,omitempty"`
	jwt.StandardClaims
}

func CreateToken(claim Claims) (tokenString string, err error) {
	claim.ExpiresAt = time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenString, err = token.SignedString([]byte("GDY"))
	return
}

func ValidateToken(signedToken string) (claims *Claims, success bool){
	token, err := jwt.ParseWithClaims(signedToken,&Claims{},secret())
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		success = true
		return
	}
	return
}

func secret()jwt.Keyfunc{
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("GDY"),nil
	}
}