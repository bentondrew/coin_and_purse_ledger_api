package main


import (
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/server"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/logger"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/router"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/db"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/api"
)


var (
  version = "0.0.28"
)


func main() {
  logger := logger.NewLogger()
  database := db.NewPostgresDatabase(logger)
  defer database.Close()
  api := api.NewAPI(database, logger)
  router := router.NewRouter(logger, "/ledger/v1.0.0", api)
  router.SetupRoutes()
  server := server.NewServer(router.Mux)
  logger.Println("Server starting.")
  err := server.ListenAndServe()
  if err != nil {
    logger.Fatalf("Server failed to start: %v", err)
  }
}
