package hello

import (
  "net/http"
  "encoding/json"
)


func HandleHello(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode("Hello World!"); err != nil {
    panic(err)
  }
}
