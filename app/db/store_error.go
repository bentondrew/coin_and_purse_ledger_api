package db

import (
    "time"
    "fmt"
)

/*StoreError is a custom error for the db package.*/
type StoreError struct {
    When time.Time
    What string
}

/*NewStoreError returns a new instance of a
StoreError*/
func NewStoreError(err string) *StoreError {
    return &StoreError{
        When: time.Now(),
        What: err,
    }
}

/*Error implements the error interface for
StoreError custom error.*/
func (se *StoreError) Error() string {
    return fmt.Sprintf("%v", se.What)
}
