package transaction

import (
  "time"
  "github.com/google/uuid"
)

/*Transaction struct is the model for transaction
instances in the ledger service. Incoming and
outgoing JSON transactions details are converted to
and from this model. This model is used to define
the database table and model instances are used
for reading and writing rows to/from the database.*/
type Transaction struct {
  ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
  Timestamp time.Time `json:"timestamp"`
  Amount float64 `json:"amount"`
}
