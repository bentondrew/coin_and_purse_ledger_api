package db


import (
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


type DataStore interface {
  InitializeDatabase() bool
  DatabaseInitialized() bool
  CreateTransaction(transaction *transaction.Transaction) error
  GetTransactions() ([]*transaction.Transaction, error)
}
