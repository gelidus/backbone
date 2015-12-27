package async

import (
	"errors"
	"sync"
)

type DoneFunc func(err error)

type SyncFunc func() error
type AsyncFunc func(done DoneFunc)

func Parallel(calls []interface{}) error {
	// create waitgroup for calls
	wait := sync.WaitGroup{}
	// error channel for the parallel service
	errchan := make(chan error, len(calls))

	for _, icall := range calls {
		// add function to the waitgroup
		wait.Add(1)

		switch call := icall.(type) {
		case SyncFunc:
			// start new synchronous parallel task
			go func() {
				err := call()
				// add error to the error channel
				if err != nil {
					errchan <- err
				}
				// after the task is done, remove it from wait group
				wait.Done()
			}()

		case AsyncFunc:
			go func() {
				call(func(err error) {
					if err != nil {
						errchan <- err
					}
					wait.Done()
				})
			}()

		default:
			return errors.New("Only SyncFunc and AsyncFunc calls are supported")
		}
	}

	// wait till all the functions are done
	wait.Wait()

	select {
	case err, ok := <- errchan:
		if !ok {
			return errors.New("Error channel was closed unexpectedly")
		} else {
			// this will only return the first error that happened
			// in the whole parallel call
			return err
		}
	default:
		// everything went well
	}

	// the end of the parallel
	return nil
}
