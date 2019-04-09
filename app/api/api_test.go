package api

import (
  "time"
  "net/http"
  "net/http/httptest"
  "testing"
  "encoding/json"
  "bytes"
  "github.com/google/uuid"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/problem"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/db"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)

func generateJSONByteArray(data interface{}) []byte {
  b, err := json.Marshal(data)
  if err != nil {
    panic(err) 
  }
  return b
}

func TestEndpointsGoodDB(t *testing.T) {
  t1, err := time.Parse(time.RFC3339, "2019-01-30T03:17:41.12004Z")
  if err != nil {
    panic(err) 
  }
  t2, err := time.Parse(time.RFC3339, "2019-01-30T19:41:10.421617Z")
  if err != nil {
    panic(err) 
  }
  id1 := uuid.New()
  id2 := uuid.New()
  reqTrans1 := `{"timestamp": "2019-01-30T03:17:41.12004Z", "amount": 10}`
  transaction1 := &transaction.Transaction{ID: id1, Timestamp: t1, Amount: 10,}
  transaction2 := &transaction.Transaction{ID: id2, Timestamp: t2, Amount: -5,}
  transactions := []*transaction.Transaction{}
  transactions = append(transactions, transaction1)
  transactions = append(transactions, transaction2) 
  mockStore := db.NewMockStore()
  mockStore.On("GetTransactions").Return(transactions, nil)
  mockStore.On("CreateTransaction").Return(transaction1, nil)
  api := NewAPI(mockStore, nil)
  tests := []struct {
    name string
    in *http.Request
    out *httptest.ResponseRecorder
    handlerFunc http.HandlerFunc
    expectedStatus int
    expectedBody string
  }{
    {
      name: "hello_get",
      in: httptest.NewRequest("GET", "/hello", nil),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleHello,
      expectedStatus: http.StatusOK,
      expectedBody: string(generateJSONByteArray("Hello World!")[:]),
    },
    {
      name: "hello_post",
      in: httptest.NewRequest("POST", "/hello", nil),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleHello,
      expectedStatus: http.StatusMethodNotAllowed,
      expectedBody: string(generateJSONByteArray(problem.Problem{Status: 405, Title: "Method Not Allowed", Detail: "Method POST is not supported by /hello", Type: "about:blank",})[:]),
    },
    {
      name: "NotFound",
      in: httptest.NewRequest("GET", "/", nil),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleDefault,
      expectedStatus: http.StatusNotFound,
      expectedBody: string(generateJSONByteArray(problem.Problem{Status: 404, Title: "Not Found", Detail: "/ not found", Type: "about:blank",})[:]),
    },
    {
      name: "transactions_get",
      in: httptest.NewRequest("GET", "/transactions", nil),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleTransactions,
      expectedStatus: http.StatusOK,
      expectedBody: string(generateJSONByteArray(transactions)[:]),
    },
    {
      name: "transactions_delete",
      in: httptest.NewRequest("DELETE", "/transactions", nil),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleTransactions,
      expectedStatus: http.StatusMethodNotAllowed,
      expectedBody: string(generateJSONByteArray(problem.Problem{Status: 405, Title: "Method Not Allowed", Detail: "Method DELETE is not supported by /transactions", Type: "about:blank",})[:]),
    },
    {
      name: "transactions_post_missing_content_type",
      in: httptest.NewRequest("POST", "/transactions", bytes.NewReader(generateJSONByteArray(reqTrans1))),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleTransactions,
      expectedStatus: http.StatusBadRequest,
      expectedBody: string(generateJSONByteArray(problem.Problem{Status: 400, Title: "Bad Request", Detail: "Field Content-Type content is missing in request header", Type: "about:blank",})[:]),
    },
    {
      name: "transactions_post_good",
      in: httptest.NewRequest("POST", "/transactions", bytes.NewReader(generateJSONByteArray(reqTrans1))),
      out: httptest.NewRecorder(),
      handlerFunc: api.HandleTransactions,
      expectedStatus: http.StatusOK,
      expectedBody: string(generateJSONByteArray(transaction1)[:]),
    },
   }
  for _, test := range tests {
    test := test
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
