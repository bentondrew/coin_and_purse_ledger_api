package main


import (
  "app/server"
  "app/logger"
  "app/router"
)


var (
  version = "0.0.6"
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
