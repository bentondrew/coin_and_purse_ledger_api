package transaction

import (
  "time"
  "github.com/google/uuid"
)


type Transaction struct {
  ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
  Timestamp time.Time `json:"timestamp"`
  Amount float64 `json:"amount"`
}
