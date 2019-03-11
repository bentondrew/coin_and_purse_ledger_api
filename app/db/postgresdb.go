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
Contains the logger for logging in the struct methods and
and a bool to indicate if the DB has been initialized
correctly.
*/
type Postgresdb struct {
  logger *log.Logger
  gormdb *gorm.DB
  initialized bool
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


func initializeDatabase(logger *log.Logger) (gormDB *gorm.DB, initialized bool) {
  defer func() {
    if rec := recover(); rec != nil {
      logger.Printf("Unable to initialize Postgres database: %s.\n", rec)
      initialized = false
      gormDB = nil
    }
  }()
  initialized = true
  gormDB = openPostgresDatabase()
  createTables(gormDB)
  return gormDB, initialized
}


/*
Returns a Postgresdb struct. Tries to open the connection
pool. If unable to, sets the initialized variable to false.
If successful in opening the db connection, tries to create
the transactions table. If unable to, sets the initialized
variable to false.
*/
func NewPostgresDatabase(logger *log.Logger) *Postgresdb {
  gormDB, initialized := initializeDatabase(logger)
  return &Postgresdb{
    logger: logger,
    gormdb: gormDB,
    initialized: initialized,
  }
}


func (p *Postgresdb) Close() {
  if err := p.gormdb.Close(); err != nil {
    panic(err)
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


func (p *Postgresdb) DatabaseInitialized() bool {
  return p.initialized
}


func (p *Postgresdb) InitializeDatabase() bool {
  gormDB, initialized := initializeDatabase(p.logger)
  p.gormdb = gormDB
  p.initialized = initialized
  return initialized
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
