package jwtoken

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/gelidus/backbone/service/crypto"
)

func Create(method crypto.Method, duration time.Duration) (string, error) {
	// create new jwt token for user that was logged in successfully
	token := jwt.New(jwt.GetSigningMethod(method.MethodName()))
	// add expiration for 3 days
	token.Claims["exp"] = time.Now().Add(duration).Unix()

	// generate string token
	return token.SignedString(method.SignKey())
}

func Valid(method crypto.Method, key string) (bool, error) {
	token, err := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		return method.VerifyKey(), nil
	})

	return token.Valid, err
}