package errors

import "fmt"

type apiError struct {
	text string
	err error
}

func NewApiError(text string, err error) *apiError {
	return &apiError{text: text, err: err}
}

func (receiver *apiError) Error() string {
	return fmt.Sprintf("error: %v", receiver.err.Error())
}

func (receiver *apiError) Unwrap() error {
	return receiver.err
}
func New(text string) string {
	return text
}