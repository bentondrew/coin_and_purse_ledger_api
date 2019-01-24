package api

import (
  "net/http"
  "encoding/json"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/db"
)


func HandleGetAllTransactions(w http.ResponseWriter, r *http.Request) {
  var transactions []transaction.Transaction
  db.Connection(func(conn gorm.DB) {
    conn.Find(&transactions)
  })
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  b, err := json.Marshal(transactions)
  if err != nil {
    panic(err) 
  } else {
    w.Write(b)
  }
}
