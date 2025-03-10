package utils

import (
	"log"
	"os"
)

// Logger instance
var Logger *log.Logger

// InitLogger initializes the logging system
func InitLogger() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	Logger = log.New(logFile, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs informational messages
func Info(msg string) {
	Logger.Println("INFO:", msg)
}

// Error logs error messages
func Error(err error) {
	Logger.Println("ERROR:", err)
}
