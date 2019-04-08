package api

import (
    "time"
    "fmt"
)

/*APIError is a custom error for the api package.*/
type APIError struct {
    When time.Time
    What string
}

/*NewAPIError returns a new instance of a
APIError*/
func NewAPIError(err string) *APIError {
    return &APIError{
        When: time.Now(),
        What: err,
    }
}

/*Error implements the error interface for
APIError custom error.*/
func (apie *APIError) Error() string {
    return fmt.Sprintf("%v", apie.What)
}
