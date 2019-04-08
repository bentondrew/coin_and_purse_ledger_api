package server


import (
  "net/http"
  "time"
)

/*NewServer returns a new instance of a custom
net/http server.*/
func NewServer(mux *http.ServeMux) *http.Server {
  return &http.Server{
    Addr: ":8080",
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout: 120 * time.Second,
    Handler: mux,
  }
}
