package sampleError

type SampleError struct {
	message string
}

func (err *SampleError) Error() string {
	return err.message
}

func New(message string) error {
	return &SampleError{
		message,
	}
}
