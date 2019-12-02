package helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClaims_GenerateValidToken(t *testing.T) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		username: "UserTest",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
		},
	}

	//Generate Token
	err, token := claims.GenerateToken("user1", time.Now().Add(1*time.Minute).Unix())
	assert.Nil(t, err)
	assert.IsType(t, "", token)
	assert.NotEmpty(t, token)

	//Validate Token
	er, isValid, username := claims.IsTokenValid(token)
	assert.Nil(t, er)
	assert.IsType(t, true, isValid)
	assert.Equal(t, true, isValid)
	assert.Equal(t, claims.username, username)
}

func TestClaims_GenerateValidTokenExpired(t *testing.T) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{}
	//Generate Token
	err, token := claims.GenerateToken("user1", time.Now().AddDate(0, 0, -1).Unix())
	assert.Nil(t, err)
	assert.IsType(t, "", token)
	assert.NotEmpty(t, token)

	//Validate Expired Token
	er, isValid, username := claims.IsTokenValid(token)
	if assert.NotNil(t, er) {
		assert.Equal(t, errors.New("Not valid token!"), er)
	}
	assert.IsType(t, false, isValid)
	assert.Equal(t, false, isValid)
	assert.Equal(t, "", username)
}
