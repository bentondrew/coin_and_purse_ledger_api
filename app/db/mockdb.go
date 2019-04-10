package db

import (
  "github.com/stretchr/testify/mock"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)

/*MockStore struct creates the mock instance for
mocking the db interaction and implements the
DataStore interface.*/
type MockStore struct {
  mock.Mock
}

/*NewMockStore returns a new instance of the
MockStore.*/
func NewMockStore() *MockStore {
  return new(MockStore)
}

/*CreateTransaction mocks the CreateTransaction method
from the DataStore interface.*/
func (ms *MockStore) CreateTransaction(transaction *transaction.Transaction) (error) {
  rets:= ms.Called()
  return rets.Error(0)
}

/*GetTransactions mocks the GetTransactions method
from the DataStore interface.*/
func (ms *MockStore) GetTransactions() ([]*transaction.Transaction, error) {
  rets := ms.Called()
  return rets.Get(0).([]*transaction.Transaction), rets.Error(1)
}
