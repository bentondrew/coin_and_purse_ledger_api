package db

import (
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


func createDbConnectString() string {
  return "postgresql://ledgerservice@roach1:26257/ledger?sslmode=disable"
}


func Connection(cb func(conn *gorm.DB)) {
  gormdb, err := gorm.Open("postgres", createDbConnectString())
  if err != nil {
    panic(err)
  }
  defer func() {
    if err := gormdb.Close(); err != nil {
      panic(err)
    }
  }()
  cb(gormdb)
}


func CreateTables() {
  Connection(func(conn *gorm.DB) {
    conn.AutoMigrate(&transaction.Transaction{})
  })
}


func NewTransactions() {
  Connection(func(conn *gorm.DB) {
    conn.Create(&transaction.Transaction{Timestamp: time.Now(), Amount: 10})
    conn.Create(&transaction.Transaction{Timestamp: time.Now(), Amount: -5})
  })
}

