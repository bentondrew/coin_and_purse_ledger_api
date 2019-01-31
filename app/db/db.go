package db

import (
  "os"
  "strings"
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


type DataStore interface {
  Find(out interface{}, where ...interface{}) *gorm.DB
  HasTable(value interface{}) bool
  AutoMigrate(values ...interface{}) *gorm.DB
  Create(value interface{}) *gorm.DB
}


func OpenPostgresDatabase() *gorm.DB {
  gormdb, err := gorm.Open("postgres", createDbConnectString())
  if err != nil {
    panic(err)
  }
  return gormdb
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


func CreateTables(db DataStore) {
  if !db.HasTable(&transaction.Transaction{}){
    db.AutoMigrate(&transaction.Transaction{})
  }
}


func NewTransactions(db DataStore) {
  db.Create(&transaction.Transaction{Timestamp: time.Now(), Amount: 10})
  db.Create(&transaction.Transaction{Timestamp: time.Now(), Amount: -5})
}
