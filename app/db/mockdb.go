package db


import (
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


type MockDatabase struct {
  Name string
}


func NewMockDatabase() *MockDatabase {
  return &MockDatabase{
    Name: "mockdb",
  }
}


func (db *MockDatabase) Find(listToPopulate *[]transaction.Transaction) {
  append(listToPopulate, transaction.Transaction{ID: 1, Timestamp: "2019-01-30T03:17:41.12004Z", Amount: 10,}, transaction.Transaction{ID: 1, Timestamp: "2019-01-30T19:41:10.421617Z", Amount: -5,})
}
