package api

import (
  "net/http"
  "encoding/json"
  "fmt"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/db"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/problem"
)


type API struct {
  store db.DataStore
}


func NewApi(store db.DataStore) *API {
  return &API {
    store: store,
  }
}


func addJsonResponseBody(data interface{}, w http.ResponseWriter) {
  b, err := json.Marshal(data)
  if err != nil {
    panic(err) 
  } else {
    w.Write(b)
  }
}


func (a *API) HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/problem+json")
  w.WriteHeader(http.StatusMethodNotAllowed)
  addJsonResponseBody(problem.Problem{Status: 405, Title: "Method Not Allowed", Detail: fmt.Sprintf("%s is not supported by %s", r.Method, r.URL), Type: "about:blank",}, w)
}


func (a *API) HandleNotFound(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/problem+json")
  w.WriteHeader(http.StatusNotFound)
  addJsonResponseBody(problem.Problem{Status: 404, Title: "Not Found", Detail: fmt.Sprintf("%s not found", r.URL), Type: "about:blank",}, w)
}


func (a *API) HandleHello(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case http.MethodGet:
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    addJsonResponseBody("Hello World!", w)
  default:
    a.HandleMethodNotAllowed(w, r)
  }
}


func (a *API) HandleGetAllTransactions(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case http.MethodGet:
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    addJsonResponseBody(a.store.GetTransactions(), w)
  default:
    a.HandleMethodNotAllowed(w, r)
  }
}
