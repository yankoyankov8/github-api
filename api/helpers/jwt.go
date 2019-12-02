package helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// Create the JWT key used to create the signature
var secretKey = []byte("my_secret_key")

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	username string `json:"username"`
	jwt.StandardClaims
}

type JwtInterface interface {
	GenerateToken(username string, expireTime int64) (error, string)
	IsTokenValid(token string) (error, bool, string)
}

func (claims *Claims) IsTokenValid(token string) (error, bool, string) {
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !tkn.Valid {
		return errors.New("Not valid token!"), false, ""
	}
	return nil, true, claims.username
}

func (claims *Claims) GenerateToken(username string, expireTime int64) (error, string) {
	//Todo think for something better
	claims.username = username
	claims.ExpiresAt = expireTime
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return errors.New("Error in creating the JWT"), ""
	}
	return nil, tokenString
}
