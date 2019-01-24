package transaction

import (
  "time"
)


type Transaction struct {
  ID int `json:"id" gorm:"primary_key"`
  Timestamp time.Time `json:"timestamp"`
  Amount float64 `json:"amount"`
}
