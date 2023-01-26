package logger

import (
	"errors"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var once = sync.Once{}
var logger *logrus.Logger

type Logger interface {
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

func Init() {
	once.Do(func() {
		logger = logrus.New()
		logger.SetReportCaller(true)
		logger.SetOutput(os.Stdout)
		logger.SetFormatter(&logrus.JSONFormatter{})
	})
}

// GetLogger returns a logger.
func GetLogger() Logger {
	if logger == nil {
		panic(errors.New("logger is not initialized"))
	}
	return logger
}
