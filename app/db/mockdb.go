package db


import (
  "log"
  "time"
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
    t1, err := time.Parse(time.RFC3339, "2019-01-30T03:17:41.12004Z")
    if err != nil {
      panic(err) 
    }
    t2, err := time.Parse(time.RFC3339, "2019-01-30T19:41:10.421617Z")
    if err != nil {
      panic(err) 
    }
    append(&out, transaction.Transaction{ID: 1, Timestamp: t1, Amount: 10,}, transaction.Transaction{ID: 1, Timestamp: t2, Amount: -5,})
  } else {
    db.logger.Println("Mockdb Find does not recognize type of provided out.")
  }
  return db
}
