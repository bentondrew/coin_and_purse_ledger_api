package main


import (
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/server"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/logger"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/router"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/db"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/api"
)


var (
  version = "0.0.20"
)


func main() {
  database := db.InitPostgresDatabase()
  defer func() {
    database.Close()
  }()
  database.CreateTables()
  database.NewTransactions()
  api := api.NewApi(database)
  logger := logger.NewLogger()
  router := router.NewRouter(logger, "/ledger/v1.0.0", api)
  router.SetupRoutes()
  server := server.NewServer(router.Mux)
  logger.Println("Server starting.")
  err := server.ListenAndServe()
  if err != nil {
    logger.Fatalf("Server failed to start: %v", err)
  }
}
