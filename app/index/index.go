package index

import (
  "net/http"
  "encoding/json"
  "fmt"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/problem"
)


func HandleIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/problem+json")
  w.WriteHeader(http.StatusNotFound)
  b, err := json.Marshal(problem.Problem{Status: 404, Title: "Not Found", Detail: fmt.Sprintf("%s not found", r.URL), Type: "about:blank",})
  if err != nil {
    panic(err) 
  } else {
    w.Write(b)
  }
}
