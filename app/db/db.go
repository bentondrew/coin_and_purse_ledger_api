package db

import (
  "time"
  "github.com/jinzhu/gorm"
    "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


var (
  dbAddr = "postgresql://ledger_service@localhost:26257/ledger?sslmode=disable"
)


func Connection(cb func(conn gorm.DB)) {
  db, err := gorm.Open("postgres", dbAddr)
  if err != nil {
    panic(err)
  }
  defer func() {
    if err := db.Close(); err != nil {
      panic(err)
    }
  }()
  cb(db)
}


func CreateTables() {
  Connection(func(conn gorm.DB) {
    conn.AutoMigrate(&transaction.Transaction{})
  })
}


func NewTransactions() {
  Connection(func(conn gorm.DB) {
    conn.Create(&Transaction{Timestamp: time.Now(), Amount: 10})
    conn.Create(&Transaction{Timestamp: time.Now(), Amount: -5})
  })
}

