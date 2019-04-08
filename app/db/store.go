package db

import (
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)

/*DataStore interface defines the functionality that the
API module needs for accessing the database. This abstraction
exists so that the database can be mocked for testing but
backed by an actual database when running the service.
Additionally, this interface can be used if in the future
it is decided to switch database implementations (like
not using an orm or changing database drivers to use
a database different than postgres).*/
type DataStore interface {
  CreateTransaction(transaction *transaction.Transaction) (error)
  GetTransactions() ([]*transaction.Transaction, error)
}
