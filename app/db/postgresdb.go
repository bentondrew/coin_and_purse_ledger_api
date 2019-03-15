package db

import (
  "os"
  "strings"
  "time"
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


/*
Struct which contains the Postgres DB connection pool.
Contains the logger for logging in the struct methods.
*/
type Postgresdb struct {
  logger *log.Logger
  gormdb *gorm.DB
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


func openPostgresDatabase() *gorm.DB {
  gormdb, err := gorm.Open("postgres", createDbConnectString())
  if err != nil {
    panic(err) 
  }
  return gormdb
}


func createTables(gormdb *gorm.DB) {
  if !gormdb.HasTable(&transaction.Transaction{}){
    if err := gormdb.AutoMigrate(&transaction.Transaction{}).Error; err != nil {
      panic(err)
    }
  }
}


func initializeDatabase(logger *log.Logger) (gormDB *gorm.DB) {
  defer func() {
    if rec := recover(); rec != nil {
      logger.Printf("Unable to initialize Postgres database: %s.\n", rec)
      gormDB = nil
    }
  }()
  gormDB = openPostgresDatabase()
  createTables(gormDB)
  return gormDB
}


/*
Returns a Postgresdb struct. Tries to open the connection
pool. If unable to, sets the initialized variable to false.
If successful in opening the db connection, tries to create
the transactions table. If unable to, sets the initialized
variable to false.
*/
func NewPostgresDatabase(logger *log.Logger) *Postgresdb {
  gormDB := initializeDatabase(logger)
  return &Postgresdb{
    logger: logger,
    gormdb: gormDB,
  }
}


func (p *Postgresdb) Close() {
  if err := p.gormdb.Close(); err != nil {
    panic(err)
  }
}


func (p *Postgresdb) DatabaseInitialized() (initialized bool) {
  if p.gormdb == nil {
    initialized = false
  } else {
    if p.gormdb.HasTable(&transaction.Transaction{}) {
      initialized = true
    } else {
      initialized = false
    }
  }
  return initialized
}


func (p *Postgresdb) NewTransactions() {
  if err := p.CreateTransaction(&transaction.Transaction{Timestamp: time.Now(), Amount: 10}); err != nil {
    panic(err) 
  }
  if err := p.CreateTransaction(&transaction.Transaction{Timestamp: time.Now(), Amount: -5}); err != nil {
    panic(err) 
  }
}


func (p *Postgresdb) checkDatabase() error {
  if p.gormdb == nil {
    p.gormdb = initializeDatabase(p.logger)
  }
  if p.gormdb != nil{
    return nil
  } else {
    return NewStoreError("Database not initialized.")
  }
}


func (p *Postgresdb) CreateTransaction(transaction *transaction.Transaction) (error) {
  err := p.checkDatabase()
  if err != nil {
    return err 
  } else {
    result := p.gormdb.Create(transaction)
    return result.Error
  }
}


func (p *Postgresdb) GetTransactions() ([]*transaction.Transaction, error) {
  err := p.checkDatabase()
  if err != nil {
    return nil, err 
  } else {
    var transactions []*transaction.Transaction
    result := p.gormdb.Find(&transactions)
    return transactions, result.Error
  }
}
