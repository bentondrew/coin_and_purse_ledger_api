package index

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "encoding/json"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/problem"
)


func TestEndpoints(t *testing.T) {
  b, err := json.Marshal(problem.Problem{Status: 404, Title: "Not Found", Detail: "/ not found", Type: "about:blank",})
  if err == nil {
    tests := []struct {
      name string
      in *http.Request
      out *httptest.ResponseRecorder
      expectedStatus int
      expectedBody string
    }{
      {
        name: "good",
        in: httptest.NewRequest("GET", "/", nil),
        out: httptest.NewRecorder(),
        expectedStatus: http.StatusNotFound,
        expectedBody: string(b[:]),
      },
     }
    for _, test := range tests {
      test := test
      t.Run(test.name, func(t *testing.T) {
        HandleIndex(test.out, test.in)
        if test.out.Code != test.expectedStatus {
          t.Logf("For Index test %s\nExpected status code: %d\nGot status code: %d\n",
                 test.name, test.expectedStatus, test.out.Code)
          t.Fail()
        }
        body := test.out.Body.String()
        if body != test.expectedBody {
          t.Logf("For Index test %s\nExpected body: %s\nGot body: %s\n",
                 test.name, test.expectedBody, body)
          t.Fail()
        }
      })
    } 
  }
}
