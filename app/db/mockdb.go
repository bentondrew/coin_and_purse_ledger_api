package db


import (
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
  "github.com/jinzhu/gorm"
)


func NewMockDatabase() *gorm.DB {
  return &gorm.DB{
    Value: "mockdb",
    Error: nil,
    RowsAffected: 0,
  }
}


func (db *gorm.DB) Find(out interface{}, where ...interface{}) *gorm.DB {
  if v, ok := out.(*[]transaction.Transaction); ok {
    append(&listToPopulate, transaction.Transaction{ID: 1, Timestamp: "2019-01-30T03:17:41.12004Z", Amount: 10,}, transaction.Transaction{ID: 1, Timestamp: "2019-01-30T19:41:10.421617Z", Amount: -5,})
  }
  return db
}
