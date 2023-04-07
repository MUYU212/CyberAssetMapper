package utils

import "fmt"

func AppendError(existError, newError error) error {
	if existError == nil {
		return newError
	}
	return fmt.Errorf("%v,%w", existError, newError)
}
