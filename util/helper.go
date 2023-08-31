// File for handling JWT tokens - when they expire
package util

import (
	"time"
	"github.com/golang-jwt/jwt"
)

const SecretKey = "manav"

//function to generate token
func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: issuer,

		//.unix() converts to unix timestamp
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // JWT will expire 24 hours after generation
	})
	return claims.SignedString([]byte(SecretKey))
}

//parses a JWT token from a given string, validates it using a secret key, and returns the issuer claim from the token's standard claims if the parsing is successful and the token is valid.
func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token)(interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}