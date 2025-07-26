package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Logger struct {
	Log *logrus.Logger
}

func NewLogger() *Logger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     false,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "msg",
		},
	})

	return &Logger{
		Log: logger,
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.Log.Infof(msg, args...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.Log.Fatalf(msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.Log.Errorf(msg, args...)
}
