package logging

import (
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init() {
	logFile, err := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Info = log.New(logFile, "INFO: ", log.Ldate|log.Ltime)
	Warning = log.New(logFile, "WARNING: ", log.Ldate|log.Ltime)
	Error = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime)
}
