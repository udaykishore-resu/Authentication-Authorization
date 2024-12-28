package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		fatalLogger: log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.infoLogger.Printf(message, args...)
}

func (l *Logger) Error(message string, args ...interface{}) {
	l.errorLogger.Printf(message, args...)
}

func (l *Logger) Fatal(message string, args ...interface{}) {
	l.fatalLogger.Fatalf(message, args...)
}
