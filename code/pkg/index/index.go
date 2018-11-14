package index

import (
  "net/http"
  "log"
  "time"
)


type Handlers struct {
  logger *log.Logger;
}


func (h *Handlers) Index(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Hello World!"))
}


func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    startTime := time.Now()
    next(w, r)
    h.logger.Printf("Request processed in %s seconds.\n", time.Now().Sub(startTime))
  }
}


func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
  mux.HandleFunc("/", h.Logger(h.Index))
}


func NewHandlers(logger *log.Logger) *Handlers {
  return &Handlers{
    logger: logger,
  }
}
