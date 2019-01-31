package db


import (
  "log"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


type DB struct {
  Value interface{}
  Error error
  RowsAffected int64
  logger *log.Logger
}


func NewMockDatabase(logger *log.Logger) *DB {
  return &DB{
    Value: "mockdb",
    Error: nil,
    RowsAffected: 0,
    logger: logger,
  }
}


func (db *DB) Find(out interface{}, where ...interface{}) *DB {
  if v, ok := out.(*[]transaction.Transaction); ok {
    append(&out, transaction.Transaction{ID: 1, Timestamp: "2019-01-30T03:17:41.12004Z", Amount: 10,}, transaction.Transaction{ID: 1, Timestamp: "2019-01-30T19:41:10.421617Z", Amount: -5,})
  } else {
    logger.Println("Mockdb Find does not recognize type of provided out.")
  }
  return db
}
