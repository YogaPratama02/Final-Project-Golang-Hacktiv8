package helpers

import (
	"BootcampHacktiv8/final_project/models"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	jwt.StandardClaims
}

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(user *models.UserLogin) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    user.Email,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Printf("Error signed string jwt with err: %s\n", err)
		return tokenString, err
	}
	return tokenString, err
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
