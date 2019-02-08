package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jackmrzhou/gc-ai/conf"
	"time"
)

type Claim struct {
	UserId uint `json:"user_id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(UserId uint, Email string) string {
	now := time.Now()
	expireTime := now.Add(48 * time.Hour)

	claim := Claim{
		UserId,
		Email,
		jwt.StandardClaims{
			ExpiresAt:expireTime.Unix(),
			Issuer:"GoogleCamp",
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, _ := tokenClaim.SignedString([]byte(conf.JWTSecret))
	return token
}

func ParseToken(token string) (*Claim, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (i interface{}, e error) {
		return conf.JWTSecret, nil
	})
	if tokenClaim != nil{
		if claim, ok := tokenClaim.Claims.(*Claim); ok && tokenClaim.Valid{
			return claim, nil
		}
	}
	return nil, err
}