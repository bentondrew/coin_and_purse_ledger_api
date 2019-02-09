package db

import (
  "os"
  "strings"
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


type Postgresdb struct {
  gormdb *gorm.DB
}


func InitPostgresDatabase() *Postgresdb {
  return &Postgresdb{
    gormdb : OpenPostgresDatabase(),
  }
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


func (p *Postgresdb) Close() {
  if err := p.gormdb.Close(); err != nil {
    panic(err)
  }
}


func (p *Postgresdb) CreateTables() {
  if !p.gormdb.HasTable(&transaction.Transaction{}){
    if err := p.gormdb.AutoMigrate(&transaction.Transaction{}).Error; err != nil {
      panic(err)
    }
  }
}


func (p *Postgresdb) NewTransactions() {
  if err := p.CreateTransaction(&transaction.Transaction{Timestamp: time.Now(), Amount: 10}); err != nil {
    panic(err) 
  }
  if err := p.CreateTransaction(&transaction.Transaction{Timestamp: time.Now(), Amount: -5}); err != nil {
    panic(err) 
  }
}


func (p *Postgresdb) CreateTransaction(transaction *transaction.Transaction) error {
  result := p.gormdb.Create(transaction)
  return result.Error
}


func (p *Postgresdb) GetTransactions() ([]*transaction.Transaction, error) {
  var transactions []*transaction.Transaction
  result := p.gormdb.Find(&transactions)
  return transactions, result.Error
}
