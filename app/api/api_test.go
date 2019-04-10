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

type testFunc func(t *testing.T)

type testValues struct {
  name string
  in *http.Request
  out *httptest.ResponseRecorder
  handlerFunc http.HandlerFunc
  expectedStatus int
  expectedBody string
}

func generateJSONByteArray(data interface{}) []byte {
  b, err := json.Marshal(data)
  if err != nil {
    panic(err) 
  }
  return b
}

func checkResults(t *testing.T,
                  w *httptest.ResponseRecorder,
                  r *http.Request,
                  expectedStatus int,
                  expectedBody string,
                  name string) {
  if w.Code != expectedStatus {
    t.Errorf("For test %s\nExpected status code: %d\nGot status code: %d\n",
             name, expectedStatus, w.Code)
  }
  body := w.Body.String()
  if body != expectedBody {
    t.Errorf("For test %s\nExpected body: %s\nGot body: %s\n",
             name, expectedBody, body)
  }
}

func TestEndpoints(t *testing.T) {
  tests := []struct {
    name string
    runFunc testFunc
  }{
    {
      name: "hello_get",
      runFunc: func(t *testing.T){
        mockStore := db.NewMockStore()
        api := NewAPI(mockStore, nil)
        values := testValues{
          name: "hello_get",
          in: httptest.NewRequest("GET", "/hello", nil),
          out: httptest.NewRecorder(),
          handlerFunc: api.HandleHello,
          expectedStatus: http.StatusOK,
          expectedBody: string(generateJSONByteArray("Hello World!")[:]),
        }
        values.handlerFunc(values.out, values.in)
        checkResults(t,
                     values.out,
                     values.in,
                     values.expectedStatus,
                     values.expectedBody,
                     values.name)
        },
    },
    {
      name: "hello_post",
      runFunc: func(t *testing.T){ 
        mockStore := db.NewMockStore()
        api := NewAPI(mockStore, nil)
        values := testValues{
          name: "hello_post",
          in: httptest.NewRequest("POST", "/hello", nil),
          out: httptest.NewRecorder(),
          handlerFunc: api.HandleHello,
          expectedStatus: http.StatusMethodNotAllowed,
          expectedBody: string(generateJSONByteArray(problem.Problem{Status: 405, Title: "Method Not Allowed", Detail: "Method POST is not supported by /hello", Type: "about:blank",})[:]),
        }
        values.handlerFunc(values.out, values.in)
        checkResults(t,
                     values.out,
                     values.in,
                     values.expectedStatus,
                     values.expectedBody,
                     values.name)
        },
    },
    {
      name: "not_found",
      runFunc: func(t *testing.T){
        mockStore := db.NewMockStore()
        api := NewAPI(mockStore, nil)
        values := testValues{
          name: "not_found",
          in: httptest.NewRequest("GET", "/", nil),
          out: httptest.NewRecorder(),
          handlerFunc: api.HandleDefault,
          expectedStatus: http.StatusNotFound,
          expectedBody: string(generateJSONByteArray(problem.Problem{Status: 404, Title: "Not Found", Detail: "/ not found", Type: "about:blank",})[:]),
        }
        values.handlerFunc(values.out, values.in)
        checkResults(t,
                     values.out,
                     values.in,
                     values.expectedStatus,
                     values.expectedBody,
                     values.name)
        },
    },
    {
      name: "transactions_get",
      runFunc: func(t *testing.T){
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
        transaction1 := &transaction.Transaction{ID: id1, Timestamp: t1, Amount: 10,}
        transaction2 := &transaction.Transaction{ID: id2, Timestamp: t2, Amount: -5,}
        transactions := []*transaction.Transaction{}
        transactions = append(transactions, transaction1)
        transactions = append(transactions, transaction2)
        mockStore := db.NewMockStore()
        mockStore.On("GetTransactions").Return(transactions, nil)
        api := NewAPI(mockStore, nil)
        values := testValues{
          name: "transactions_get",
          in: httptest.NewRequest("GET", "/transactions", nil),
          out: httptest.NewRecorder(),
          handlerFunc: api.HandleTransactions,
          expectedStatus: http.StatusOK,
          expectedBody: string(generateJSONByteArray(transactions)[:]),
        }
        values.handlerFunc(values.out, values.in)
        checkResults(t,
                     values.out,
                     values.in,
                     values.expectedStatus,
                     values.expectedBody,
                     values.name)
        },
    },
    {
      name: "transactions_delete",
      runFunc: func(t *testing.T){
        mockStore := db.NewMockStore()
        api := NewAPI(mockStore, nil)
        values := testValues{
          name: "transactions_delete",
          in: httptest.NewRequest("DELETE", "/transactions", nil),
          out: httptest.NewRecorder(),
          handlerFunc: api.HandleTransactions,
          expectedStatus: http.StatusMethodNotAllowed,
          expectedBody: string(generateJSONByteArray(problem.Problem{Status: 405, Title: "Method Not Allowed", Detail: "Method DELETE is not supported by /transactions", Type: "about:blank",})[:]),
        }
        values.handlerFunc(values.out, values.in)
        checkResults(t,
                     values.out,
                     values.in,
                     values.expectedStatus,
                     values.expectedBody,
                     values.name)
        },
    },
    {
      name: "transactions_post_missing_content_type",
      runFunc: func(t *testing.T){
        reqTrans1 := `{"timestamp": "2019-01-30T03:17:41.12004Z", "amount": 10}`
        mockStore := db.NewMockStore()
        api := NewAPI(mockStore, nil)
        values := testValues{
          name: "transactions_post_missing_content_type",
          in: httptest.NewRequest("POST", "/transactions", bytes.NewReader(generateJSONByteArray(reqTrans1))),
          out: httptest.NewRecorder(),
          handlerFunc: api.HandleTransactions,
          expectedStatus: http.StatusBadRequest,
          expectedBody: string(generateJSONByteArray(problem.Problem{Status: 400, Title: "Bad Request", Detail: "Field Content-Type is missing in request header", Type: "about:blank",})[:]),
        }
        values.handlerFunc(values.out, values.in)
        checkResults(t,
                     values.out,
                     values.in,
                     values.expectedStatus,
                     values.expectedBody,
                     values.name)
        },
    },
    {
      name: "transactions_post_missing_content_type_value",
      runFunc: func(t *testing.T){
        reqTrans1 := `{"timestamp": "2019-01-30T03:17:41.12004Z", "amount": 10}`
        mockStore := db.NewMockStore()
        api := NewAPI(mockStore, nil)
        mockRequest := httptest.NewRequest("POST", "/transactions", bytes.NewReader(generateJSONByteArray(reqTrans1)))
        nilString := nil
        mockRequest.Header.Set("Content-Type", &nilString)
        values := testValues{
          name: "transactions_post_missing_content_type_value",
          in: mockRequest,
          out: httptest.NewRecorder(),
          handlerFunc: api.HandleTransactions,
          expectedStatus: http.StatusBadRequest,
          expectedBody: string(generateJSONByteArray(problem.Problem{Status: 400, Title: "Bad Request", Detail: "Field Content-Type content is empty in request header", Type: "about:blank",})[:]),
        }
        values.handlerFunc(values.out, values.in)
        checkResults(t,
                     values.out,
                     values.in,
                     values.expectedStatus,
                     values.expectedBody,
                     values.name)
        },
    },
    {
      name: "transactions_post_good",
      runFunc: func(t *testing.T){
        t1, err := time.Parse(time.RFC3339, "2019-01-30T03:17:41.12004Z")
        if err != nil {
          panic(err) 
        }
        id1 := uuid.New()
        reqTrans1 := `{"timestamp": "2019-01-30T03:17:41.12004Z", "amount": 10}`
        transaction1 := &transaction.Transaction{ID: id1, Timestamp: t1, Amount: 10,}
        mockStore := db.NewMockStore()
        mockStore.On("CreateTransaction").Return(transaction1, nil)
        api := NewAPI(mockStore, nil)
        mockRequest := httptest.NewRequest("POST", "/transactions", bytes.NewReader(generateJSONByteArray(reqTrans1)))
        mockRequest.Header.Set("Content-Type", "application/json")
        values := testValues{
          name: "transactions_post_good",
          in: mockRequest,
          out: httptest.NewRecorder(),
          handlerFunc: api.HandleTransactions,
          expectedStatus: http.StatusOK,
          expectedBody: string(generateJSONByteArray(transaction1)[:]),
        }
        values.handlerFunc(values.out, values.in)
        checkResults(t,
                     values.out,
                     values.in,
                     values.expectedStatus,
                     values.expectedBody,
                     values.name)
        },
    },
   }
  for _, test := range tests {
    test := test
    t.Run(test.name, test.runFunc)
  }
}
