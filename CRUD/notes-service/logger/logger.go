package logger

import (
	"log"
	"runtime"
)

// LogWithCaller logs a message with information about the caller
func LogWithCaller(message string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Println("Failed to get caller information")
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	log.Printf("[file: %s, line: %d, func: %s] %s", file, line, funcName, message)
}
