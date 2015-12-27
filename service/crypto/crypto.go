package crypto

type Method interface {
	MethodName() string
	SignKey() interface{}
	VerifyKey() interface{}
}
