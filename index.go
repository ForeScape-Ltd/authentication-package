package authenticationpkg

//package main

import (
	"fmt"
	"time"

	"github.com/ForeScape-Ltd/authentication-package/token"
)

func TODO(tokenSecretKey, refreshTokenSecretKey string, tokenDuration, refreshTokenDuration time.Duration) (token.Maker, error) {

	jwt, err := token.NewJWTMaker(tokenSecretKey, refreshTokenSecretKey, tokenDuration, refreshTokenDuration)
	return jwt, err
}

func main() {
	token, err := TODO("12345678901234567890123456789012", "12345678901234567890123456789012", time.Minute, time.Minute)
	fmt.Println("Hello Gift")
	if err != nil {
		fmt.Println(err.Error())
	}

	tok, pay, err := token.CreateToken("admin@gmail.com")
	fmt.Println("Hello Mirabel")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Bearer " + tok)
	fmt.Println(pay)
}
