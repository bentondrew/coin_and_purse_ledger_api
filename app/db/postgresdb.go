package db

import (
	"log"
	"os"
	"strings"
	"github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*Postgresdb struct contains the Postgres DB connection pool and
the logger for logging in the struct methods. This struct also
implements the DataStore interface.*/
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
	if !gormdb.HasTable(&transaction.Transaction{}) {
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

/*NewPostgresDatabase returns a Postgresdb struct.
Tries to open the connection pool. If unable to, sets
the initialized variable to false. If successful in opening
the db connection, tries to create the transactions table.
If unable to, sets the initialized variable to false.*/
func NewPostgresDatabase(logger *log.Logger) *Postgresdb {
	gormDB := initializeDatabase(logger)
	return &Postgresdb{
		logger: logger,
		gormdb: gormDB,
	}
}

/*Close tries to close the postgres db pool.
Panics if error.*/
func (p *Postgresdb) Close() {
	if err := p.gormdb.Close(); err != nil {
		panic(err)
	}
}

func (p *Postgresdb) checkDatabase() error {
	if p.gormdb == nil {
		p.gormdb = initializeDatabase(p.logger)
	}
	if p.gormdb != nil {
		return nil
	}
	return NewStoreError("Database not initialized.")
}

/*CreateTransaction tries to add the provided transaction struct
to the database. Returns any errors. Error assumes transaction was
not added. This is the implementation of the DataStore interface
CreateTransaction.*/
func (p *Postgresdb) CreateTransaction(transaction *transaction.Transaction) (err error) {
	// To catch panic from uuid New
	defer func() {
		if rec := recover(); rec != nil {
			err = rec.(error)
		}
	}()
	err = p.checkDatabase()
	if err == nil {
		id := uuid.New()
		transaction.ID = id
		result := p.gormdb.Create(transaction)
		err = result.Error
	}
	return err
}

/*GetTransactions returns a slice of all the transaction structs in the database.
If successful, the error return should be nil. If not successful the slice
should be nil and error populated. This is the implementation of the DataStore
interface GetTransactions.*/
func (p *Postgresdb) GetTransactions() ([]*transaction.Transaction, error) {
	err := p.checkDatabase()
	if err != nil {
		return nil, err
	}
	var transactions []*transaction.Transaction
	result := p.gormdb.Find(&transactions)
	return transactions, result.Error
}
