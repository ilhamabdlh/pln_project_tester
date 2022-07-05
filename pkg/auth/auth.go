package auth

import (
	"errors"
	"time"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Name string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateJWT(name string, password string) (tokenString string, err error){
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Name: name,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
func ValidateToken(signedToken string) (err error){
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error){
			return []byte(jwtKey), nil
		},
	)
	fmt.Println("token: ", string(jwtKey))
	if err != nil{
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok{
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix(){
		err = errors.New("token expired")
		return
	}
	return
}