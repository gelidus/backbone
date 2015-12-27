package crypto

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
)

type RSA struct {
	private *rsa.PrivateKey
	public  *rsa.PublicKey
}

func NewRSAMethod(privateKey, publicKey []byte) (*RSA, error) {
	private, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}

	public, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, err
	}

	rsa := &RSA{
		private,
		public,
	}

	return rsa, nil
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