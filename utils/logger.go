package utils

import (
    "log"
    "os"
)

// Logger instance
var Logger = log.New(os.Stdout, "[OpenStack-Service] ", log.Ldate|log.Ltime|log.Lshortfile)

// LogError logs errors with a specific prefix
func LogError(err error) {
    if err != nil {
        Logger.Println("ERROR:", err)
    }
}

// LogInfo logs general information
func LogInfo(msg string) {
    Logger.Println("INFO:", msg)
}
