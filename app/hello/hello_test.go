package hello

import (
  "net/http"
  "net/http/httptest"
  "testing"
)


func TestEndpoints(t *testing.T) {
  tests := []struct {
    name string
    in *http.Request
    out *httptest.ResponseRecorder
    expectedStatus int
    expectedBody string
  }{
    {
      name: "good",
      in: httptest.NewRequest("GET", "/test", nil),
      out: httptest.NewRecorder(),
      expectedStatus: http.StatusOK,
      expectedBody: "Hello World!",
    },
   }
  for _, test := range tests {
    test := test
    t.Run(test.name, func(t *testing.T) {
      HandleHello(test.out, test.in)
      if test.out.Code != test.expectedStatus {
        t.Logf("For Hello test %s\nExpected status code: %d\nGot status code: %d\n",
               test.name, test.expectedStatus, test.out.Code)
        t.Fail()
      }
      body := test.out.Body.String()
      if body != test.expectedBody {
        t.Logf("For Hello test %s\nExpected body: %s\nGot body: %s\n",
               test.name, test.expectedBody, body)
      }
    })
  }
}
