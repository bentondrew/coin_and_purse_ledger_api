package api

import (
  "net/http"
  "encoding/json"
  "fmt"
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
  }
  w.Write(b)
}


func (api *API) errorRecovery(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, req *http.Request) {
    defer func() {
      if rec := recover(); rec != nil {
        api.HandleServerError(w, req, rec)
      }
    }()
    next(w, req)
  }
}


func (api *API) HandleServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
  w.Header().Set("Content-Type", "application/problem+json")
  w.WriteHeader(http.StatusInternalServerError)
  addJsonResponseBody(problem.Problem{Status: 500, Title: "Internal Server Error", Detail: fmt.Println(err), Type: "about:blank",}, w)
}


func (api *API) HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/problem+json")
  w.WriteHeader(http.StatusMethodNotAllowed)
  addJsonResponseBody(problem.Problem{Status: 405, Title: "Method Not Allowed", Detail: fmt.Sprintf("%s is not supported by %s", r.Method, r.URL), Type: "about:blank",}, w)
}


func (api *API) HandleNotFound(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/problem+json")
  w.WriteHeader(http.StatusNotFound)
  addJsonResponseBody(problem.Problem{Status: 404, Title: "Not Found", Detail: fmt.Sprintf("%s not found", r.URL), Type: "about:blank",}, w)
}


func (api *API) HandleDefault(w http.ResponseWriter, r *http.Request) {
  api.errorRecovery(api.HandleNotFound)(w, r)
}


func (api *API) HandleHello(w http.ResponseWriter, r *http.Request) {
  api.errorRecovery(api.helloResponseGeneration)(w, r)
}


func (api *API) helloResponseGeneration(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case http.MethodGet:
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    addJsonResponseBody("Hello World!", w)
  default:
    api.HandleMethodNotAllowed(w, r)
  }
}


func (api *API) HandleGetAllTransactions(w http.ResponseWriter, r *http.Request) {
  api.errorRecovery(api.getAllTransactionsResponseGeneration)(w, r)
}


func (api *API) getAllTransactionsResponseGeneration(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case http.MethodGet:
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    transactions, err := a.store.GetTransactions()
    if err != nil {
      panic(err) 
    }
    addJsonResponseBody(transactions, w)
  default:
    api.HandleMethodNotAllowed(w, r)
  }
}
