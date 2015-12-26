package core

type Validable interface {
	Valid() (bool, error)
}
