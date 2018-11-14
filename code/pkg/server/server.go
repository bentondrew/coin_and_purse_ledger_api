package server

import (
  "net/http"
  "time"
)


func New(mux *http.ServeMux) *http.Server {
  srv := &http.Server{
    Addr: ":8080",
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout: 120 * time.Second,
    Handler: mux,
  }
  return srv
}
