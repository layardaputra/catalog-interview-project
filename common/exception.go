package common

import "fmt"

// Define custom error types
type CustomError struct {
	StatusCode int
	Message    string
	Err        error
}

// Error implements the error interface for CustomError.
func (e *CustomError) Error() string {
	return fmt.Sprintf("CustomError: %s, Original Error: %v", e.Message, e.Err.Error())
}
