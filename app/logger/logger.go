package logger

import (
  "log"
  "os"
)


func NewLogger() *log.Logger{
  return log.New(os.Stdout, "ledger_service ", log.LstdFlags|log.Lshortfile)
}
