package api

import (
  "net/http"
  "encoding/json"
  "fmt"
  "log"
  "io"
  "io/ioutil"
  "time"
  "strings"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/db"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/problem"
)

/*API struct stores the data store used for the API endpoints.
Also implements the methods for the API endpoints.*/
type API struct {
  store db.DataStore
  logger *log.Logger
}

type responseGenerator func(w http.ResponseWriter, r *http.Request) (int, []byte)

/*NewAPI takes the provided data store and returns
an initialized API struct with the data
store populated.*/
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
            /*handleServerError uses json marshal. This second error catch manually overrides so that
            there is a definite exit point.*/
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

/*responseWriter should be used right before any response is sent
to write the desired header and body to the response.*/
func (api *API) responseWriter(w http.ResponseWriter, statusCode int, body []byte) {
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

/*HandleDefault is used for returning a general resource not found (404) or a internal
service error (500) if there is an issue processing the 404.*/
func (api *API) HandleDefault(w http.ResponseWriter, r *http.Request) {
  api.apiHandlerFunc(api.handleNotFound)(w,r)
}

func (api *API) handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) (int, []byte) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 405, Title: "Method Not Allowed", Detail: fmt.Sprintf("Method %s is not supported by %s", r.Method, r.URL), Type: "about:blank",})
  if err != nil {
    panic(err) 
  }
  return http.StatusMethodNotAllowed, b
}

func (api *API) handleBadRequestMissingHeaderField(w http.ResponseWriter, r *http.Request, mf string) (int, []byte) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 400, Title: "Bad Request", Detail: fmt.Sprintf("Field %s is missing in request header", mf), Type: "about:blank",})
  if err != nil {
    panic(err) 
  }
  return http.StatusBadRequest, b
}

func (api *API) handleBadRequestMissingHeaderFieldContent(w http.ResponseWriter, r *http.Request, field string) (int, []byte) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 400, Title: "Bad Request", Detail: fmt.Sprintf("Field %s content is empty in request header", field), Type: "about:blank",})
  if err != nil {
    panic(err) 
  }
  return http.StatusBadRequest, b
}

func (api *API) handleUnsupportedMediaType(w http.ResponseWriter, r *http.Request) (int, []byte) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 415, Title: "Unsupported Media Type", Detail: fmt.Sprintf("Content type %s is not supported by %s", r.Header.Get("Content-Type"), r.URL), Type: "about:blank",})
  if err != nil {
    panic(err) 
  }
  return http.StatusUnsupportedMediaType, b
}

func (api *API) handleRequestTooLarge(w http.ResponseWriter, r *http.Request) (int, []byte) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 413, Title: "Request Entity Too Large", Detail: fmt.Sprintf("Request sent to %s is too large.", r.URL), Type: "about:blank",})
  if err != nil {
    panic(err) 
  }
  return http.StatusRequestEntityTooLarge, b
}

func (api *API) handleMissingObjectKey(w http.ResponseWriter, r *http.Request, missingKey string) (int, []byte) {
  w.Header().Set("Content-Type", "application/problem+json")
  b, err := json.Marshal(problem.Problem{Status: 422, Title: "Unprocessable Entity", Detail: fmt.Sprintf("Request JSON object sent to %s is missing key: '%s'.", r.URL, missingKey), Type: "about:blank",})
  if err != nil {
    panic(err) 
  }
  return http.StatusUnprocessableEntity, b
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

/*HandleHello is a hello world test endpoint.*/
func (api *API) HandleHello(w http.ResponseWriter, r *http.Request) {
  api.apiHandlerFunc(api.helloResponseGeneration)(w, r)
}

func (api *API) transactionGet(w http.ResponseWriter, r *http.Request) (status int, b []byte) {
  params := r.URL.Query()
  if len(params) == 0 {
    transactions, err := api.store.GetTransactions()
    if err != nil {
      panic(err) 
    }
    jsonBytes, err := json.Marshal(transactions)
    if err != nil {
      panic(err) 
    }
    b = jsonBytes
    status = http.StatusOK
  } else
  {
    outputString := "Query contents: \n"
    for k, v := range params {
      keyString := fmt.Sprintf("Key: %s, Value: %s\n", k, v)
      outputString = outputString + keyString
    }
    api.logger.Println(outputString)
    jsonBytes, err := json.Marshal(outputString)
    if err != nil {
      panic(err) 
    }
    b = jsonBytes
    status = http.StatusOK
  }
  w.Header().Set("Content-Type", "application/json")
  return status, b
}

func (api *API) transactionPost(w http.ResponseWriter, r *http.Request) (status int, b []byte) {
  /*
  TODO:
    Need to write tests for this function.
  */
  contentTypeHeaderKey := "Content-Type"
  contentType, ok := r.Header[contentTypeHeaderKey]
  if ok {
    if len(contentType) > 0 {
      /*Currently ignores any options in the content field.*/
      if contentType[0] == "application/json" {
        reqBody, errR := ioutil.ReadAll(io.LimitReader(r.Body, 524288000))
        if errR != nil {
          /*Assumes failure due to request size. May need to make
          more granular in the future.*/
          return api.handleRequestTooLarge(w, r) 
        }
        if err := r.Body.Close(); err != nil {
          panic(err)
        }
        var reqObj map[string]interface{}
        if errUm := json.Unmarshal(reqBody, &reqObj); errUm != nil {
          panic(errUm)
        }
        var transaction transaction.Transaction
        foundTime := false
        foundAmount := false
        for k, v := range reqObj {
          if strings.ToLower(k) == "timestamp" {
            tmstmp, err := time.Parse(time.RFC3339, v.(string))
            if err != nil {
              panic(err) 
            }
            transaction.Timestamp = tmstmp
            foundTime = true
          }
          if strings.ToLower(k) == "amount" {
            transaction.Amount = v.(float64)
            foundAmount = true
          }
        }
        if !foundTime {
          return api.handleMissingObjectKey(w, r, "timestamp")
        }
        if !foundAmount {
          return api.handleMissingObjectKey(w, r, "amount")
        }
        if errCt := api.store.CreateTransaction(&transaction); errCt != nil {
          panic(errCt)
        }
        jsonBytes, errJm := json.Marshal(transaction)
        if errJm != nil {
          panic(errJm)
        }
        b = jsonBytes
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Location", fmt.Sprintf("/transactions/%d", transaction.ID))
        return http.StatusCreated, b
      }
      return api.handleUnsupportedMediaType(w, r)
    }
    return api.handleBadRequestMissingHeaderFieldContent(w, r, contentTypeHeaderKey)
  }
  return api.handleBadRequestMissingHeaderField(w, r, contentTypeHeaderKey)
}

func (api *API) transactionsResponseGeneration(w http.ResponseWriter, r *http.Request) (status int, b []byte) {
  switch r.Method {
  case http.MethodGet:
    return api.transactionGet(w, r)
  case http.MethodPost:
    return api.transactionPost(w, r)
  default:
    return api.handleMethodNotAllowed(w, r)
  }
}

/*HandleTransactions is the endpoint for handling transaction requests.
Posts require JSON objects.*/
func (api *API) HandleTransactions(w http.ResponseWriter, r *http.Request) {
  api.apiHandlerFunc(api.transactionsResponseGeneration)(w, r)
}
