package async

import (
	"errors"
	"time"
)

type AsynchronousResult interface {
	Done() bool
}

func WaitForResult[T AsynchronousResult](timeOutSeconds int, checkResult func() (T, error)) (T, error) {
	timeOutAt := time.Now().Add(time.Second * time.Duration(timeOutSeconds))

	var zero T
	for {
		if time.Now().After(timeOutAt) {
			return zero, errors.New("result timed out")
		}

		result, err := checkResult()
		if err != nil {
			return zero, err
		}

		if result.Done() {
			return result, nil
		}

		time.Sleep(time.Millisecond * 10)
	}
}
