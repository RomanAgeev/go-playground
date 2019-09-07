package sampleError

import "fmt"

func InvalidParamCount(expected int, actual int) error {
	return fmt.Errorf("Invalid param count: expected %d, but got %d", expected, actual)
}
