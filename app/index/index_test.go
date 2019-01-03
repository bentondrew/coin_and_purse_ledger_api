package index

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "fmt"
  "encoding/json"
  "app/problem"
)


func TestEndpoints(t *testing.T) {
  tests := []struct {
    name string
    in *http.Request
    out *httptest.ResponseRecorder
    expectedStatus int
    expectedBody []byte
  }{
    {
      name: "good",
      in: httptest.NewRequest("GET", "/", nil),
      out: httptest.NewRecorder(),
      expectedStatus: http.StatusNotFound,
      expectedBody: json.Marshal(problem.Problem{Status: 404, Title: "Not Found", Detail: fmt.Sprintf("%s not found", r.URL), Type: "about:blank"}),
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
      body := test.out.Body
      if body != test.expectedBody {
        t.Logf("For Index test %s\nExpected body: %s\nGot body: %s\n",
               test.name, test.expectedBody, body)
      }
    })
  }
}
