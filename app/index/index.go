package index

import (
  "net/http"
  "json"
  "fmt"
  "app/problem"
)


func HandleIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json+problem")
  w.WriteHeader(http.StatusNotFound)
  json.NewEncoder(w).Encode(problem.Problem{Status: 404, Title: "Not Found", Detail: fmt.Sprintf("%s not found", r.URL), Type: "about:blank"})
}
