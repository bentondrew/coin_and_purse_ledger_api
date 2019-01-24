package api

import (
  "net/http"
  "encoding/json"
  "fmt"
  "github.com/jinzhu/gorm"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/db"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/problem"
)


func HandleNotFound(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/problem+json")
  w.WriteHeader(http.StatusNotFound)
  b, err := json.Marshal(problem.Problem{Status: 404, Title: "Not Found", Detail: fmt.Sprintf("%s not found", r.URL), Type: "about:blank",})
  if err != nil {
    panic(err) 
  } else {
    w.Write(b)
  }
}


func HandleHello(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.WriteHeader(http.StatusOK)
  b, err := json.Marshal("Hello World!")
  if err != nil {
    panic(err) 
  } else {
    w.Write(b)
  }
}


func HandleGetAllTransactions(w http.ResponseWriter, r *http.Request) {
  var transactions []transaction.Transaction
  db.Connection(func(conn *gorm.DB) {
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
