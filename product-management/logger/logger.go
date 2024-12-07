package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Global logger instance
var log *logrus.Logger

// Init initializes the logger with structured logging
func Init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})
}

// Info logs an info level message
func Info(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

// Error logs an error level message
func Error(msg string, err error) {
	log.Errorf(msg+" : %v", err)
}
