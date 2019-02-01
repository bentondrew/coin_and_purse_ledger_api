package db


import (
  "log"
  "time"
  "reflect"
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
  switch reflect.TypeOf(out).Kind() {
  case reflect.Slice:
    switch reflect.TypeOf(out).Elem().Kind() {
    case reflect.Struct:
      switch {
      case reflect.TypeOf(out).Elem().Name() == transaction.Transaction:
        t1, err := time.Parse(time.RFC3339, "2019-01-30T03:17:41.12004Z")
        if err != nil {
          panic(err) 
        }
        t2, err := time.Parse(time.RFC3339, "2019-01-30T19:41:10.421617Z")
        if err != nil {
          panic(err) 
        }
        append(reflect.TypeOf(out).Elem(), transaction.Transaction{ID: 1, Timestamp: t1, Amount: 10,}, transaction.Transaction{ID: 1, Timestamp: t2, Amount: -5,})
      default:
        db.logger.Println("Mockdb Find currently only handles struct of type transaction.Transaction.")
      }
    default:
      db.logger.Println("Mockdb Find currently only handles struct in slice.")
    }
  default:
    db.logger.Println("Mockdb Find currently only handles slice.")
  }
  return db
}
