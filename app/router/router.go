package router

import (
  "net/http"
  "log"
  "time"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/hello"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/index"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/transaction"
)


type Router struct {
  Mux *http.ServeMux
  logger *log.Logger
} 


func NewRouter(logger *log.Logger) *Router {
  return &Router{
    Mux: http.NewServeMux(),
    logger: logger,
  }
}


func (rtr *Router) EndpointLogger(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    startTime := time.Now()
    next(w, r)
    rtr.logger.Printf("Request to path %s processed in %s seconds.\n", r.URL, time.Now().Sub(startTime))
  }
}


func (rtr *Router) SetupRoutes() {
  rtr.Mux.HandleFunc("/", rtr.EndpointLogger(index.HandleIndex))
  rtr.Mux.HandleFunc("/hello", rtr.EndpointLogger(hello.HandleHello))
  rtr.Mux.HandleFunc("/transactions", rtr.EndpointLogger(transaction.HandleGetAllTransactions))
}
