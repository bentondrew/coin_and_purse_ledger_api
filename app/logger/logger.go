package logger

import (
  "log"
  "os"
)

/*NewLogger returns an instance of a logger configured for this service.*/
func NewLogger() *log.Logger{
  return log.New(os.Stdout, "ledger_service ", log.LstdFlags|log.Lshortfile)
}
