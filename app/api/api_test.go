package api

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "encoding/json"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/problem"
)


func TestEndpoints(t *testing.T) {
  b_h, err := json.Marshal("Hello World!")
  if err != nil {
    panic(err) 
  }
  b_nf, err := json.Marshal(problem.Problem{Status: 404, Title: "Not Found", Detail: "/ not found", Type: "about:blank",})
  if err != nil {
    panic(err) 
  }
  tests := []struct {
    name string
    in *http.Request
    out *httptest.ResponseRecorder
    handlerFunc http.HandlerFunc
    expectedStatus int
    expectedBody string
  }{
    {
      name: "hello",
      in: httptest.NewRequest("GET", "/hello", nil),
      out: httptest.NewRecorder(),
      handlerFunc: HandleHello,
      expectedStatus: http.StatusOK,
      expectedBody: string(b_h[:]),
    },
    {
      name: "NotFound",
      in: httptest.NewRequest("GET", "/", nil),
      out: httptest.NewRecorder(),
      handlerFunc: HandleNotFound,
      expectedStatus: http.StatusNotFound,
      expectedBody: string(b_nf[:]),
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
