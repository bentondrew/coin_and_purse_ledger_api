package api


import (
    "time"
    "fmt"
)


type APIError struct {
    When time.Time
    What string
}


func NewAPIError(err string) *APIError {
    return &APIError{
        When: time.Now(),
        What: err,
    }
}


func (apie *APIError) Error() string {
    return fmt.Sprintf("%v: %v", apie.When, apie.What)
}
