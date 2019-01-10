package transaction

import (
  "time"
  "net/http"
  "encoding/json"
)


type Transaction struct {
  Timestamp time.Time `json: "timestamp"`
  Amount float64 `json: "amount"`
}


type Transactions []Transaction


var(
  transactions = Transactions{
    Transaction{Timestamp: time.Now(), Amount: 10},
    Transaction{Timestamp: time.Now(), Amount: -5},
  }
)


func HandleGetAllTransactions(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(transactions); err != nil {
    panic(err)
  }
}
