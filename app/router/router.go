package router

import (
  "net/http"
  "log"
  "time"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/api"
)


type Router struct {
  Mux *http.ServeMux
  logger *log.Logger
  baseUrl string
} 


func NewRouter(logger *log.Logger, baseUrl string) *Router {
  return &Router{
    Mux: http.NewServeMux(),
    logger: logger,
    baseUrl: baseUrl,
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
  rtr.Mux.HandleFunc("/", rtr.EndpointLogger(api.HandleNotFound))
  rtr.Mux.HandleFunc(rtr.baseUrl + "/hello", rtr.EndpointLogger(api.HandleHello))
  rtr.Mux.HandleFunc(rtr.baseUrl + "/transactions", rtr.EndpointLogger(api.HandleGetAllTransactions))
}
