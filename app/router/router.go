package router

import (
  "net/http"
  "log"
  "time"
  "app/hello"
  "app/index"
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
    rtr.logger.Printf("Request processed in %s seconds.\n", time.Now().Sub(startTime))
  }
}


func (rtr *Router) SetupRoutes() {
  rtr.Mux.HandleFunc("/", rtr.EndpointLogger(index.HandleIndex))
  rtr.Mux.HandleFunc("/test", rtr.EndpointLogger(hello.HandleHello))
}
