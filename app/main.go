package main


import (
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/server"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/logger"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/router"
)


var (
  version = "0.0.13"
)


func main() {
  logger := logger.NewLogger()
  router := router.NewRouter(logger)
  router.SetupRoutes()
  server := server.NewServer(router.Mux)
  logger.Println("Server starting.")
  err := server.ListenAndServe()
  if err != nil {
    logger.Fatalf("Server failed to start: %v", err)
  }
}
