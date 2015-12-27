package crypto

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
)

type RSA struct {
	public  *rsa.PublicKey
	private *rsa.PrivateKey
}

func NewRSAMethod(privateKey, publicKey []byte) *RSA {
	rsa := &RSA{
		jwt.ParseRSAPrivateKeyFromPEM(privateKey),
		jwt.ParseRSAPublicKeyFromPEM(publicKey),
	}

	return rsa
}

func (method *RSA) MethodName() string {
	return "RS256"
}

func (method *RSA) SignKey() interface{} {
	return method.private
}

func (method *RSA) VerifyKey() interface{} {
	return method.public
}