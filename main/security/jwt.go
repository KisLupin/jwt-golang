package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Main function
func main() {
	a := generateToken()
	fmt.Println(a)
	b, err := validateToken(a)
	fmt.Println(b, err)
}

var mySigningKey = []byte("mysupersecretkey")

type jwtCustomClaims struct {
	Name string `json:"name"`
	Admin bool  `json:"admin"`
	jwt.StandardClaims
}

func generateToken() string {
	claims := jwtCustomClaims{
		"linh", true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    "bad",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}
	return t
}

func validateToken(token string) (*jwt.Token, error){
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
}
