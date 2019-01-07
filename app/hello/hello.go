package hello

import (
  "net/http"
  "encoding/json"
)


func HandleHello(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode("Hello World!")
}
