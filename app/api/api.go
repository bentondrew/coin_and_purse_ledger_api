package api

import (
  "net/http"
  "encoding/json"
  "fmt"
  "log"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/db"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/problem"
)


/*
API struct stores the data store used for the API endpoints.
Also implements the methods for the API endpoints.
*/
type API struct {
  store db.DataStore
  logger *log.Logger
}


type responseGenerator func(w http.ResponseWriter, r *http.Request) (int, []byte)


/*
Takes the provided data store and returns
an initialized API struct with the data
store populated.
*/
func NewAPI(store db.DataStore, logger *log.Logger) *API {
  return &API {
    store: store,
    logger: logger,
  }
}


func (api *API) handleServerError(w http.ResponseWriter, r *http.Request, err interface{}) (int, []byte) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 500, Title: "Internal Server Error", Detail: fmt.Sprintf("%s", err), Type: "about:blank",})
  if err != nil {
    panic(err) 
  }
  return http.StatusInternalServerError, b
}


func (api *API) errorRecovery(next responseGenerator) responseGenerator {
  return func(w http.ResponseWriter, req *http.Request) (statusCode int, body []byte) {
    defer func() {
      if rec := recover(); rec != nil {
        defer func() {
          if rec := recover(); rec != nil {
            // handleServerError uses json marshal. This second error catch manually overrides so that
            // there is a definite exit point.
            w.Header().Set("Content-Type", "application/problem+json")
            errorString := fmt.Sprintf("%s", rec)
            jsonString := `{"status": 500, "title": "Internal Server Error", "detail": "` + errorString + `", "type": "about:blank"}`
            body = []byte(jsonString)
            statusCode = http.StatusInternalServerError
          }
        }()
        statusCode, body = api.handleServerError(w, req, rec)
      }
    }()
    statusCode, body = next(w, req)
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
    api.responseWriter(w, statusCode, body) 
  }
}


func (api *API) handleNotFound(w http.ResponseWriter, r *http.Request) (int, []byte) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 404, Title: "Not Found", Detail: fmt.Sprintf("%s not found", r.URL), Type: "about:blank",})
  if err != nil {
    panic(err) 
  }
  return http.StatusNotFound, b
}


func (api *API) HandleDefault(w http.ResponseWriter, r *http.Request) {
  api.apiHandlerFunc(api.handleNotFound)(w,r)
}


func (api *API) handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) (int, []byte) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 405, Title: "Method Not Allowed", Detail: fmt.Sprintf("%s is not supported by %s", r.Method, r.URL), Type: "about:blank",})
  if err != nil {
    panic(err) 
  }
  return http.StatusMethodNotAllowed, b
}


func (api *API) helloResponseGeneration(w http.ResponseWriter, r *http.Request) (int, []byte) {
  switch r.Method {
  case http.MethodGet:
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    b, err := json.Marshal("Hello World!")
    if err != nil {
      panic(err) 
    }
    return http.StatusOK, b
  default:
    return api.handleMethodNotAllowed(w, r)
  }
}


func (api *API) HandleHello(w http.ResponseWriter, r *http.Request) {
  api.apiHandlerFunc(api.helloResponseGeneration)(w, r)
}


func (api *API) databaseInitialized() (initialized bool) {
  initialized = api.store.DatabaseInitialized()
  if !initialized {
    initialized = api.store.InitializeDatabase()
  }
  return initialized
}


func (api *API) transactionsResponseGeneration(w http.ResponseWriter, r *http.Request) (status int, b []byte) {
  switch r.Method {
  case http.MethodGet:
    if api.databaseInitialized(){
      params := r.URL.Query()
      if len(params) == 0 {
        transactions, err := api.store.GetTransactions()
        if err != nil {
          panic(err) 
        }
        b, err := json.Marshal(transactions)
        if err != nil {
          panic(err) 
        }
        status = http.StatusOK
      } else
      {
        outputString := "Query contents: \n"
        for k, v := range params {
          keyString := fmt.Sprintf("Key: %s, Value: %s\n", k, v)
          outputString = outputString + keyString
        }
        api.logger.Println(outputString)
      }
      w.Header().Set("Content-Type", "application/json")
      return status, b
    } else {
      panic(NewAPIError("Cannot get transactions; database not initialized."))
    }
  default:
    return api.handleMethodNotAllowed(w, r)
  }
}


func (api *API) HandleTransactions(w http.ResponseWriter, r *http.Request) {
  api.apiHandlerFunc(api.transactionsResponseGeneration)(w, r)
}
