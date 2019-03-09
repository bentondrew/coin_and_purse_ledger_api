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


type responseGenerator func(w http.ResponseWriter, r *http.Request) (int, []byte, error)


func NewApi(store db.DataStore) *API {
  return &API {
    store: store,
  }
}


func (api *API) handleServerError(w http.ResponseWriter, r *http.Request, err interface{}) (int, []byte, error) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 500, Title: "Internal Server Error", Detail: fmt.Sprintf("%s", err), Type: "about:blank",})
  return http.StatusInternalServerError, b, err
}


func (api *API) errorRecovery(next responseGenerator) (int, []byte) {
  return func(w http.ResponseWriter, req *http.Request) (int, []byte) {
    statusCode, body, err := next(w, req)
    if err != nil {
      statusCode, body, err := api.handleServerError(w, req, err)
      if err != nil {
        // handleServerError uses json marshal, this second error catching manually overrides so that
        // there is a definite exit point.
        error_string := fmt.Sprintf("%s", rec)
        json_string := `{"status": 500, "title": "Internal Server Error", "detail": "` + error_string + `", "type": "about:blank"}`
        body := []byte(json_string)
        statusCode := http.StatusInternalServerError
      }
    }
    return statusCode, body
  }
}


func (api *API) responseWriter(w http.ResponseWriter, statusCode int, body []byte) {
  /*
  This function should be used right before any response is sent
  to write the desired header and body to the response.
  */
  w.WriteHeader(statusCode)
  w.Write(body)
}


func (api *API) apiHandlerFunc(next responseGenerator) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    statusCode, body := api.errorRecovery(next)(w, r)
    responseWriter(w, statusCode, body) 
  }
}


func (api *API) handleNotFound(w http.ResponseWriter, r *http.Request) (int, []byte, error) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 404, Title: "Not Found", Detail: fmt.Sprintf("%s not found", r.URL), Type: "about:blank",})
  return http.StatusNotFound, b, err
}


func (api *API) HandleDefault(w http.ResponseWriter, r *http.Request) {
  api.apiHandlerFunc(api.handleNotFound)(w,r)
}


func (api *API) handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) (int, []byte, error) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 405, Title: "Method Not Allowed", Detail: fmt.Sprintf("%s is not supported by %s", r.Method, r.URL), Type: "about:blank",})
  return http.StatusMethodNotAllowed, b, err
}


func (api *API) helloResponseGeneration(w http.ResponseWriter, r *http.Request) (int, []byte, error) {
  switch r.Method {
  case http.MethodGet:
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    b, err := json.Marshal("Hello World!")
    return http.StatusOK, b, err
  default:
    return api.handleMethodNotAllowed(w, r)
  }
}


func (api *API) HandleHello(w http.ResponseWriter, r *http.Request) {
  api.apiHandlerFunc(api.helloResponseGeneration)(w, r)
}


func (api *API) getAllTransactionsResponseGeneration(w http.ResponseWriter, r *http.Request) (int, []byte, error) {
  switch r.Method {
  case http.MethodGet:
    transactions, err := api.store.GetTransactions()
    if err != nil {
      return nil, nil, err 
    }
    w.Header().Set("Content-Type", "application/json")
    b, err := json.Marshal(transactions)
    return http.StatusOK, b, err
  default:
    return api.handleMethodNotAllowed(w, r)
  }
}


func (api *API) HandleGetAllTransactions(w http.ResponseWriter, r *http.Request) {
  api.apiHandlerFunc(api.getAllTransactionsResponseGeneration)(w, r)
}
