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


func (db *DB) HasTable(value interface{}) bool {
  db.logger.Println("Mockdb HasTable not implemented.")
  return false
}


func (db *DB) AutoMigrate(values ...interface{}) *DB {
  db.logger.Println("Mockdb AutoMigrate not implemented.")
  return db
}


func (db *DB) Create(value interface{}) *DB {
  db.logger.Println("Mockdb Create not implemented.")
  return db
}


func (db *DB) Find(out interface{}, where ...interface{}) *DB {
  switch reflect.TypeOf(out).Kind() {
  case reflect.Slice:
    switch reflect.TypeOf(out).Elem().Kind() {
    case reflect.Struct:
      switch {
      case reflect.TypeOf(out).Elem().Name() == "transaction.Transaction":
        t1, err := time.Parse(time.RFC3339, "2019-01-30T03:17:41.12004Z")
        if err != nil {
          panic(err) 
        }
        t2, err := time.Parse(time.RFC3339, "2019-01-30T19:41:10.421617Z")
        if err != nil {
          panic(err) 
        }
        sl := reflect.ValueOf(out).Elem()
        sl.Set(reflect.Append(sl, reflect.ValueOf(transaction.Transaction{ID: 1, Timestamp: t1, Amount: 10,}), reflect.ValueOf(transaction.Transaction{ID: 1, Timestamp: t2, Amount: -5,})))
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
