package db


import (
  "github.com/stretchr/testify/mock"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


type MockStore struct {
  mock.Mock
}


func NewMockStore() *MockStore {
  return new(MockStore)
}


func (ms *MockStore) CreateTransaction(transaction *transaction.Transaction) (*transaction.Transaction, error) {
  rets:= ms.Called(transaction)
  return rets.Get(0).(*transaction.Transaction), rets.Error(1)
}


func (ms *MockStore) GetTransactions() ([]*transaction.Transaction, error) {
  rets := ms.Called()
  return rets.Get(0).([]*transaction.Transaction), rets.Error(1)
}
