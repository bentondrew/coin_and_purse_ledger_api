package db

import (
  "os"
  "strings"
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


type Database struct {
  ConnectionPool *gorm.DB
}


func NewDatabase() *Database {
  gormdb, err := gorm.Open("postgres", createDbConnectString())
  if err != nil {
    panic(err)
  }
  return &Database{
    ConnectionPool: gormdb,
  }
}


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


func (db *Database) CreateTables() {
  if !db.ConnectionPool.HasTable(&transaction.Transaction{}){
    db.ConnectionPool.AutoMigrate(&transaction.Transaction{})
  }
}


func (db *Database) NewTransactions() {
  db.ConnectionPool.Create(&transaction.Transaction{Timestamp: time.Now(), Amount: 10})
  db.ConnectionPool.Create(&transaction.Transaction{Timestamp: time.Now(), Amount: -5})
}

