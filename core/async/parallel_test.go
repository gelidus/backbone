package async

import (
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)

func TestParallelSync(t *testing.T) {
	x := 0
	// just nothing
	func1 := func() error {
		return nil
	}
	// should assign 5 to upper scope
	func2 := func() error {
		x = 5
		return nil
	}
	// should return error
	func3 := func() error {
		return errors.New("Some error")
	}

	err := Parallel([]interface{}{
		SyncFunc(func1), SyncFunc(func2),
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, 5, x)

	err = Parallel([]interface{}{
		SyncFunc(func1), SyncFunc(func2), SyncFunc(func3),
	})

	assert.NotEqual(t, nil, err)
}

func TestParallelAsync(t *testing.T) {
	// correct function
	func1 := func(done DoneFunc) {
		done(nil)
	}
	// test
	err := Parallel([]interface{}{
		AsyncFunc(func1),
	})

	assert.Equal(t, nil, err)

	// error function
	func2 := func(done DoneFunc) {
		done(errors.New("Some error"))
	}
	// test
	err = Parallel([]interface{}{
		AsyncFunc(func1), AsyncFunc(func2),
	})

	assert.NotEqual(t, nil, err)
}

func TestParallelMixed(t *testing.T) {
	func1 := func() error {
		return nil
	}

	func2 := func(done DoneFunc) {
		done(nil)
	}

	err := Parallel([]interface{}{
		SyncFunc(func1), AsyncFunc(func2),
	})

	assert.Equal(t, nil, err)
}