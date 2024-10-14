package main

import (
	"fmt"
)

type Logger interface {
	Info(message string)
	Error(message string)
}

type ThirdPartyLogger struct{}

func (l *ThirdPartyLogger) Log(level string, message string) {
	fmt.Printf("[%s]: %s\n", level, message)
}

type LoggerAdapter struct {
	thirdPartyLogger *ThirdPartyLogger
}

func (a *LoggerAdapter) Info(message string) {
	a.thirdPartyLogger.Log("INFO", message)
}

func (a *LoggerAdapter) Error(message string) {
	a.thirdPartyLogger.Log("ERROR", message)
}

func Process(logger Logger) {
	logger.Info("Process is start")
	logger.Error("Error in process")
}

func main() {
	thirdPartyLogger := &ThirdPartyLogger{}
	loggerAdapter := &LoggerAdapter{
		thirdPartyLogger: thirdPartyLogger,
	}

	Process(loggerAdapter)
}
