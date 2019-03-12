package db

import (
    "time"
    "fmt"
)


type StoreError struct {
    When time.Time
    What string
}


func NewStoreError(err string) *StoreError {
    return &StoreError{
        When: time.Now(),
        What: err,
    }
}


func (se *StoreError) Error() string {
    return fmt.Sprintf("%v", se.What)
}
