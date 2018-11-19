package main


import (
  "net/http"
  "log"
  "os"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/code/pkg/index"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/code/pkg/server"
)


var (
  version = "0.0.4"
)


func main() {
  logger := log.New(os.Stdout, "ledger_service ", log.LstdFlags|log.Lshortfile)
  idx := index.NewHandlers(logger)
  mux := http.NewServeMux()
  idx.SetupRoutes(mux)
  srv := server.New(mux)
  logger.Println("Server starting.")
  err := srv.ListenAndServe()
  if err != nil {
    logger.Fatalf("Server failed to start: %v", err)
  }
}
