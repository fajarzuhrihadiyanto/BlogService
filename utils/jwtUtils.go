package utils

import (
	"github.com/golang-jwt/jwt"
)

// CreateToken
// This function is used to wrap some data into json web token
func CreateToken(payload jwt.MapClaims) (string, error) {
	// Get the secret key
	secret := GetEnvVariable("JWT_SECRET_KEY")

	// Configure JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Sign the token with the secret key and return the token string
	return token.SignedString([]byte(secret))
}

// VerifyToken
// This function is used to verify given token string and extract the data inside
func VerifyToken(token string) (jwt.MapClaims, jwt.Token, error) {
	// Get the secret key
	secret := GetEnvVariable("JWT_SECRET_KEY")

	// Parse the token string
	claims := &jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return *claims, *tkn, err
}
