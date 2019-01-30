package db

import (
  "os"
  "strings"
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


func createDbConnectString() string {
  dbUser := ""
  dbPass := ""
  dbHost := ""
  dbPort := ""
  dbName := ""
  for _, e := range os.Environ() {
    pair := strings.Split(e, "=")
    switch {
    case pair[0] == "DB_HOST":
      dbHost = pair[1]
    case pair[0] == "DB_PORT":
      dbPort = pair[1]
    case pair[0] == "DB_USER":
      dbUser = pair[1]
    case pair[0] == "DB_PASS":
      dbPass = pair[1]
    case pair[0] == "DB_DATABASE":
      dbName = pair[1]
    }
  }
  return "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
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
    if !conn.HasTable(&transaction.Transaction{}){
      conn.AutoMigrate(&transaction.Transaction{})
    }
  })
}


func NewTransactions() {
  Connection(func(conn *gorm.DB) {
    conn.Create(&transaction.Transaction{Timestamp: time.Now(), Amount: 10})
    conn.Create(&transaction.Transaction{Timestamp: time.Now(), Amount: -5})
  })
}

