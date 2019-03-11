package api

import (
  "time"
  "net/http"
  "net/http/httptest"
  "testing"
  "encoding/json"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/problem"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/db"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


func generateJsonByteArray(data interface{}) []byte {
  b, err := json.Marshal(data)
  if err != nil {
    panic(err) 
  }
  return b
}


func TestEndpoints(t *testing.T) {
  t1, err := time.Parse(time.RFC3339, "2019-01-30T03:17:41.12004Z")
  if err != nil {
    panic(err) 
  }
  t2, err := time.Parse(time.RFC3339, "2019-01-30T19:41:10.421617Z")
  if err != nil {
    panic(err) 
  }
  transactions := []*transaction.Transaction{}
  transactions = append(transactions, &transaction.Transaction{ID: 1, Timestamp: t1, Amount: 10,})
  transactions = append(transactions, &transaction.Transaction{ID: 2, Timestamp: t2, Amount: -5,}) 
  tests := []struct {
    name string
    in *http.Request
    out *httptest.ResponseRecorder
    handlerFunc http.HandlerFunc
    expectedStatus int
    expectedBody string
    mockTransactions []*transaction.Transaction
    mockGetTransactionsError error
    mockDBInitialize bool
    mockDBInitializeCheck bool
  }{
    {
      name: "hello_get",
      in: httptest.NewRequest("GET", "/hello", nil),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleHello,
      expectedStatus: http.StatusOK,
      expectedBody: string(generateJsonByteArray("Hello World!")[:]),
      mockTransactions: transactions,
      mockGetTransactionsError: nil,
      mockDBInitialize: true,
      mockDBInitializeCheck: true,
    },
    {
      name: "hello_post",
      in: httptest.NewRequest("POST", "/hello", nil),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleHello,
      expectedStatus: http.StatusMethodNotAllowed,
      expectedBody: string(generateJsonByteArray(problem.Problem{Status: 405, Title: "Method Not Allowed", Detail: "POST is not supported by /hello", Type: "about:blank",})[:]),
      mockTransactions: transactions,
      mockGetTransactionsError: nil,
      mockDBInitialize: true,
      mockDBInitializeCheck: true,
    },
    {
      name: "NotFound",
      in: httptest.NewRequest("GET", "/", nil),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleDefault,
      expectedStatus: http.StatusNotFound,
      expectedBody: string(generateJsonByteArray(problem.Problem{Status: 404, Title: "Not Found", Detail: "/ not found", Type: "about:blank",})[:]),
      mockTransactions: transactions,
      mockGetTransactionsError: nil,
      mockDBInitialize: true,
      mockDBInitializeCheck: true,
    },
    {
      name: "transactions_get",
      in: httptest.NewRequest("GET", "/transactions", nil),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleGetAllTransactions,
      expectedStatus: http.StatusOK,
      expectedBody: string(generateJsonByteArray(transactions)[:]),
      mockTransactions: transactions,
      mockGetTransactionsError: nil,
      mockDBInitialize: true,
      mockDBInitializeCheck: true,
    },
    {
      name: "transactions_post",
      in: httptest.NewRequest("POST", "/transactions", nil),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleGetAllTransactions,
      expectedStatus: http.StatusMethodNotAllowed,
      expectedBody: string(generateJsonByteArray(problem.Problem{Status: 405, Title: "Method Not Allowed", Detail: "POST is not supported by /transactions", Type: "about:blank",})[:]),
      mockTransactions: transactions,
      mockGetTransactionsError: nil,
      mockDBInitialize: true,
      mockDBInitializeCheck: true,
    },
   }
  for _, test := range tests {
    test := test
    mockStore := db.NewMockStore()
    mockStore.On("GetTransactions").Return(test.mockTransactions, test.mockGetTransactionsError).Once()
    mockStore.On("InitializeDatabase").Return(test.mockDBInitialize).Once()
    mockStore.On("DatabaseInitialized").Return(test.mockDBInitializeCheck).Once()
    api := NewAPI(mockStore)
    t.Run(test.name, func(t *testing.T) {
      test.handlerFunc(test.out, test.in)
      if test.out.Code != test.expectedStatus {
        t.Logf("For test %s\nExpected status code: %d\nGot status code: %d\n",
               test.name, test.expectedStatus, test.out.Code)
        t.Fail()
      }
      body := test.out.Body.String()
      if body != test.expectedBody {
        t.Logf("For test %s\nExpected body: %s\nGot body: %s\n",
               test.name, test.expectedBody, body)
        t.Fail()
      }
    })
  } 
}
