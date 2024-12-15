package utils

import "errors"

// Common Errors
var (
	ErrNotFound         = errors.New("Resource not found")
	ErrInvalidInput     = errors.New("Invalid input")
	ErrDatabaseError    = errors.New("Database error")
	ErrProcessingFailed = errors.New("Processing failed")
)

// HandleError is a utility function for error handling
func HandleError(err error, message string) error {
	if err != nil {
		return errors.New(message + ": " + err.Error())
	}
	return nil
}
