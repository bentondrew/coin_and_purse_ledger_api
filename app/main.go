package main


import (
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/server"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/logger"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/router"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/db"
)


var (
  version = "0.0.18"
)


func main() {
  db.CreateTables()
  db.NewTransactions()
  logger := logger.NewLogger()
  router := router.NewRouter(logger, "/ledger/v1.0.0")
  router.SetupRoutes()
  server := server.NewServer(router.Mux)
  logger.Println("Server starting.")
  err := server.ListenAndServe()
  if err != nil {
    logger.Fatalf("Server failed to start: %v", err)
  }
}
