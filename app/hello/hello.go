package hello

import (
  "net/http"
)


func HandleHello(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Hello World!"))
}
